// Copyright (c) nano Author and TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/nats-io/nuid"
	"go.opentelemetry.io/otel/attribute"

	"ricebean/pkg/acceptor"
	"ricebean/pkg/pipeline"

	"ricebean/pkg/agent"
	"ricebean/pkg/cluster"
	"ricebean/pkg/component"
	"ricebean/pkg/conn/codec"
	"ricebean/pkg/conn/message"
	"ricebean/pkg/conn/packet"
	"ricebean/pkg/constants"
	pcontext "ricebean/pkg/context"
	"ricebean/pkg/docgenerator"
	e "ricebean/pkg/errors"
	"ricebean/pkg/logger"
	"ricebean/pkg/metrics"
	"ricebean/pkg/route"
	"ricebean/pkg/serialize"
	"ricebean/pkg/timer"
	"ricebean/pkg/tracing"
)

var (
	handlerType = "handler"
)

type (
	// HandlerService service
	HandlerService struct {
		baseService
		chLocalProcess   chan unhandledMessage // channel of messages that will be processed locally
		chRemoteProcess  chan unhandledMessage // channel of messages that will be processed remotely
		packetCodec      codec.PacketCodec     // binary decoder
		remoteService    *RemoteService
		serializer       serialize.Serializer          // message serializer
		server           *cluster.Server               // server obj
		services         map[string]*component.Service // all registered service
		metricsReporters []metrics.Reporter
		agentFactory     agent.AgentFactory
		handlerPool      *HandlerPool
		handlers         map[string]*component.Handler // all handler method
		messageCodec     message.MessageCodec

		_OnGatewayReceive func(ctx context.Context, a agent.Agent, route *route.Route, msg *message.Message) //接收返回未处理的消息数据
	}

	unhandledMessage struct {
		ctx   context.Context
		agent agent.Agent
		route *route.Route
		msg   *message.Message
	}
)

// GatewayHandlerService NewHandlerService creates and returns a new handler service
func NewHandlerService(
	packetCodec codec.PacketCodec,
	messageCodec message.MessageCodec,
	serializer serialize.Serializer,
	localProcessBufferSize int,
	remoteProcessBufferSize int,
	server *cluster.Server,
	remoteService *RemoteService,
	agentFactory agent.AgentFactory,
	metricsReporters []metrics.Reporter,
	handlerHooks *pipeline.HandlerHooks,
	handlerPool *HandlerPool,
) *HandlerService {
	h := &HandlerService{
		services:         make(map[string]*component.Service),
		chLocalProcess:   make(chan unhandledMessage, localProcessBufferSize),
		chRemoteProcess:  make(chan unhandledMessage, remoteProcessBufferSize),
		packetCodec:      packetCodec,
		messageCodec:     messageCodec,
		serializer:       serializer,
		server:           server,
		remoteService:    remoteService,
		agentFactory:     agentFactory,
		metricsReporters: metricsReporters,
		handlerPool:      handlerPool,
		handlers:         make(map[string]*component.Handler),
	}

	h.handlerHooks = handlerHooks

	return h
}

// Dispatch message to corresponding logic handler 将消息发送给相应的逻辑处理程序
func (h *HandlerService) Dispatch(thread int) {
	// TODO: This timer is being stopped multiple times, it probably doesn't need to be stopped here
	defer timer.GlobalTicker.Stop()

	for {
		// Calls to remote servers block calls to local server
		select {
		case lm := <-h.chLocalProcess:
			metrics.ReportMessageProcessDelayFromCtx(lm.ctx, h.metricsReporters, "local")
			h.localProcess(lm.ctx, lm.agent, lm.route, lm.msg)

		case rm := <-h.chRemoteProcess:
			metrics.ReportMessageProcessDelayFromCtx(rm.ctx, h.metricsReporters, "remote")
			h.remoteService.remoteProcess(rm.ctx, nil, rm.agent, rm.route, rm.msg)

		case <-timer.GlobalTicker.C: // execute cron task
			timer.Cron()

		case t := <-timer.Manager.ChCreatedTimer: // new Timers
			timer.AddTimer(t)

		case id := <-timer.Manager.ChClosingTimer: // closing Timers
			timer.RemoveTimer(id)
		}
	}
}

// Register registers components
func (h *HandlerService) Register(comp component.Component, opts []component.Option) error {
	s := component.NewService(comp, opts)

	if _, ok := h.services[s.Name]; ok {
		return fmt.Errorf("handler: service already defined: %s", s.Name)
	}

	if err := s.ExtractHandler(); err != nil {
		return err
	}

	// register all handlers
	h.services[s.Name] = s
	for name, handler := range s.Handlers {
		h.handlerPool.Register(s.Name, name, handler)
	}
	return nil
}

// Handle handles messages from a conn
func (h *HandlerService) Handle(conn acceptor.PlayerConn) {
	// create a client agent and startup write goroutine
	a := h.agentFactory.CreateAgent(conn)

	// startup agent goroutine
	go a.Handle()

	logger.Log.Debugf("New session established: %s", a.String())

	// guarantee agent related resource is destroyed
	defer func() {
		a.GetSession().Close()
		logger.Log.Debugf("Session read goroutine exit, SessionID=%d, UID=%s", a.GetSession().ID(), a.GetSession().UID())
	}()

	for {
		msg, err := conn.GetNextMessage()

		if err != nil {
			// Check if this is an expected error due to connection being closed
			if errors.Is(err, net.ErrClosed) {
				logger.Log.Debugf("Connection no longer available while reading next available message: %s", err.Error())
			} else if err == constants.ErrConnectionClosed {
				logger.Log.Debugf("Connection no longer available while reading next available message: %s", err.Error())
			} else {
				// Differentiate errors for valid sessions, to avoid noise from load balancer healthchecks and other internet noise
				if a.GetStatus() != constants.StatusStart {
					logger.Log.Errorf("Error reading next available message for UID: %s, Build: %s, error: %s", a.GetSession().UID(), "a.GetSession().GetHandshakeData().Sys.BuildNumber", err.Error())
				} else {
					logger.Log.Debugf("Error reading next available message on initial connection: %s", err.Error())
				}
			}

			return
		}

		packets, err := h.packetCodec.Decode(msg)
		if err != nil {
			logger.Log.Errorf("Failed to decode message: %s", err.Error())
			return
		}

		if len(packets) < 1 {
			logger.Log.Warnf("Read no packets, data: %v", msg)
			continue
		}

		// process all packet
		for i := range packets {
			if err := h.processPacket(a, packets[i]); err != nil {
				logger.Log.Errorf("Failed to process packet: %s", err.Error())
				return
			}
		}
	}
}

func (h *HandlerService) processPacket(a agent.Agent, p *packet.Packet) error {
	/*
		switch p.Type {
		case packet.Handshake: //握手
			logger.Log.Debug("Received handshake packet")

			// Parse the json sent with the handshake by the client
			handshakeData := &session.HandshakeData{}
			if err := json.Unmarshal(p.Data, handshakeData); err != nil {
				defer a.Close()
				logger.Log.Errorf("Failed to unmarshal handshake data: %s", err.Error())
				if serr := a.SendHandshakeErrorResponse(); serr != nil {
					logger.Log.Errorf("Error sending handshake error response: %s", err.Error())
					return err
				}

				return fmt.Errorf("invalid handshake data. Id=%d", a.GetSession().ID())
			}

			if err := a.GetSession().ValidateHandshake(handshakeData); err != nil {
				defer a.Close()
				logger.Log.Errorf("Handshake validation failed: %s", err.Error())
				if serr := a.SendHandshakeErrorResponse(); serr != nil {
					logger.Log.Errorf("Error sending handshake error response: %s", err.Error())
					return err
				}

				return fmt.Errorf("handshake validation failed: %w. SessionId=%d", err, a.GetSession().ID())
			}
			//发送握手响应
			if err := a.SendHandshakeResponse(); err != nil {
				logger.Log.Errorf("Error sending handshake response: %s", err.Error())
				return err
			}
			logger.Log.Debugf("Session handshake Id=%d, Remote=%s", a.GetSession().ID(), a.RemoteAddr())

			a.GetSession().SetHandshakeData(handshakeData)
			a.SetStatus(constants.StatusHandshake)
			err := a.GetSession().Set(constants.IPVersionKey, a.IPVersion())
			if err != nil {
				logger.Log.Warnf("failed to save ip version on session: %q\n", err)
			}

			logger.Log.Debug("Successfully saved handshake data")

		case packet.HandshakeAck: //握手确认
			a.SetStatus(constants.StatusWorking)
			logger.Log.Debugf("Receive handshake ACK Id=%d, Remote=%s", a.GetSession().ID(), a.RemoteAddr())

		case packet.Data: //数据
			if a.GetStatus() < constants.StatusWorking {
				return fmt.Errorf("receive data on socket which is not yet ACK, session will be closed immediately, remote=%s",
					a.RemoteAddr().String())
			}

			msg, err := message.Decode(p.Data)
			if err != nil {
				return err
			}
			h.processMessage(a, msg)

		case packet.Heartbeat:
			// expected
		}
	*/

	a.SetStatus(constants.StatusWorking)
	if a.GetStatus() < constants.StatusWorking {
		return fmt.Errorf("receive data on socket which is not yet ACK, session will be closed immediately, remote=%s",
			a.RemoteAddr().String())
	}

	msg, err := h.messageCodec.Decode(p.Data)
	if err != nil {
		return err
	}
	msg.Route = ""
	h.processMessage(a, msg)

	a.SetLastAt() //设置最后时间
	return nil
}

// 设置接收客户端消息处理
func (h *HandlerService) SetOnGatewayReceive(callback func(ctx context.Context, a agent.Agent, route *route.Route, msg *message.Message)) {
	h._OnGatewayReceive = callback
}

// 处理消息. b
func (h *HandlerService) processMessage(a agent.Agent, msg *message.Message) {
	requestID := nuid.New().Next()
	ctx := pcontext.AddToPropagateCtx(context.Background(), constants.StartTimeKey, time.Now().UnixNano())
	ctx = pcontext.AddToPropagateCtx(ctx, constants.RouteKey, msg.Route)
	ctx = pcontext.AddToPropagateCtx(ctx, constants.RequestIDKey, requestID)

	//属性
	attributes := []attribute.KeyValue{
		attribute.String("local.id", h.server.ID),
		attribute.String("span.kind", "server"),
		attribute.String("msg.type", strings.ToLower(msg.Type.String())),
		attribute.String("user.id", a.GetSession().UID()),
		attribute.String("request.id", requestID),
	}
	ctx, _ = tracing.StartSpan(ctx, msg.Route, attributes...)
	ctx = context.WithValue(ctx, constants.SessionCtxKey, a.GetSession())

	r, err := route.Decode(msg.Route)
	if err != nil {
		logger.Log.Errorf("Failed to decode route: %s", err.Error())
		a.AnswerWithError(ctx, msg.ID, e.NewError(err, e.ErrBadRequestCode))
		return
	}

	if r.SvType == "" {
		r.SvType = h.server.Type
	}

	message := unhandledMessage{
		ctx:   ctx,
		agent: a,
		route: r,
		msg:   msg,
	}
	if r.SvType == h.server.Type {
		h.chLocalProcess <- message
	} else {
		if h.remoteService != nil {
			h.chRemoteProcess <- message
		} else {
			logger.Log.Warnf("request made to another server type but no remoteService running")
		}
	}
}

func (h *HandlerService) localProcess(ctx context.Context, a agent.Agent, route *route.Route, msg *message.Message) {
	/*
		var mid uint
		switch msg.Type {
		case message.Request:
			mid = msg.ID
		case message.Notify:
			mid = 0
		}
	*/

	h._OnGatewayReceive(ctx, a, route, msg)

	/*
		ret, err := h.handlerPool.ProcessHandlerMessage(ctx, route, h.serializer, h.handlerHooks, a.GetSession(), msg.Data, msg.Type, false)
		if msg.Type != message.Notify {
			if err != nil {
				logger.Log.Errorf("Failed to process handler message: %s", err.Error())
				a.AnswerWithError(ctx, mid, err)
			} else {
				err := a.GetSession().ResponseMID(ctx, mid, ret)
				if err != nil {
					logger.Log.Errorf("Failed to process handler message: %s", err.Error())
					tracing.FinishSpan(ctx, err)
					metrics.ReportTimingFromCtx(ctx, h.metricsReporters, handlerType, err)
				}
			}
		} else {
			//来自 Ctx 的报告时间
			metrics.ReportTimingFromCtx(ctx, h.metricsReporters, handlerType, err)
			tracing.FinishSpan(ctx, err)
			if err != nil {
				logger.Log.Errorf("Failed to process notify message: %s", err.Error())
			}
		}
	*/
}

// DumpServices outputs all registered services
func (h *HandlerService) DumpServices() {
	handlers := h.handlerPool.GetHandlers()
	for name := range handlers {
		logger.Log.Infof("registered handler %s, isRawArg: %v", name, handlers[name].IsRawArg)
	}
}

// Docs returns documentation for handlers
func (h *HandlerService) Docs(getPtrNames bool) (map[string]interface{}, error) {
	if h == nil {
		return map[string]interface{}{}, nil
	}
	return docgenerator.HandlersDocs(h.server.Type, h.services, getPtrNames)
}

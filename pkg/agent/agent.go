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

package agent

import (
	"context"
	gojson "encoding/json"
	e "errors"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"ricebean/pkg/config"
	"ricebean/pkg/conn/codec"
	"ricebean/pkg/conn/message"
	"ricebean/pkg/conn/packet"
	"ricebean/pkg/constants"
	"ricebean/pkg/errors"
	"ricebean/pkg/logger"
	"ricebean/pkg/metrics"
	"ricebean/pkg/protos"
	"ricebean/pkg/serialize"
	"ricebean/pkg/session"
	"ricebean/pkg/tracing"
	"ricebean/pkg/util"
	"ricebean/pkg/util/compression"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var (
	// hbd contains the heartbeat packet data
	hbd []byte
	// hrd contains the handshake response data
	hrd []byte
	// herd contains the handshake error response data
	herd []byte
	once sync.Once
)

const handlerType = "handler"

type (
	agentImpl struct {
		Session            session.Session // session
		sessionPool        session.SessionPool
		appDieChan         chan bool         // app die channel
		chDie              chan struct{}     // wait for close
		chSend             chan pendingWrite // push message queue
		chStopHeartbeat    chan struct{}     // stop heartbeats
		chStopWrite        chan struct{}     // stop writing messages
		closeMutex         sync.Mutex
		conn               net.Conn // low-level conn fd
		packetCodec        codec.PacketCodec
		heartbeatTimeout   time.Duration
		writeTimeout       time.Duration
		lastAt             int64 // last heartbeat unix time stamp
		messageCodec       message.MessageCodec
		messagesBufferSize int // size of the pending messages buffer
		metricsReporters   []metrics.Reporter
		serializer         serialize.Serializer // message serializer
		state              int32                // current agent state
	}

	pendingMessage struct {
		ctx     context.Context
		typ     message.Type // message type
		route   string       // message route (push)
		mid     uint         // response message id (response)
		payload interface{}  // payload
		err     bool         // if its an error message
	}

	pendingWrite struct {
		ctx  context.Context
		data []byte
		err  error
	}

	// Agent corresponds to a user and is used for storing raw Conn information
	Agent interface {
		GetSession() session.Session
		Push(route string, v interface{}) error
		ResponseMID(ctx context.Context, mid uint, v interface{}, isError ...bool) error
		Close() error
		RemoteAddr() net.Addr
		String() string
		GetStatus() int32
		Kick(ctx context.Context) error
		SetLastAt()
		SetStatus(state int32)
		Handle()
		IPVersion() string
		SendHandshakeResponse() error
		SendHandshakeErrorResponse() error
		SendRequest(ctx context.Context, serverID, route string, v interface{}) (*protos.Response, error)
		AnswerWithError(ctx context.Context, mid uint, err error)
	}

	// AgentFactory factory for creating Agent instances
	AgentFactory interface {
		CreateAgent(conn net.Conn) Agent
	}

	agentFactoryImpl struct {
		sessionPool        session.SessionPool
		appDieChan         chan bool // app die channel
		packetCodec        codec.PacketCodec
		heartbeatTimeout   time.Duration
		writeTimeout       time.Duration
		messageCodec       message.MessageCodec
		messagesBufferSize int // size of the pending messages buffer
		metricsReporters   []metrics.Reporter
		serializer         serialize.Serializer // message serializer
	}
)

// NewAgentFactory ctor
func NewAgentFactory(
	appDieChan chan bool,
	packetCodec codec.PacketCodec,
	serializer serialize.Serializer,
	heartbeatTimeout time.Duration,
	writeTimeout time.Duration,
	messageCodec message.MessageCodec,
	messagesBufferSize int,
	sessionPool session.SessionPool,
	metricsReporters []metrics.Reporter,
) AgentFactory {
	return &agentFactoryImpl{
		appDieChan:         appDieChan,
		packetCodec:        packetCodec,
		heartbeatTimeout:   heartbeatTimeout,
		writeTimeout:       writeTimeout,
		messageCodec:       messageCodec,
		messagesBufferSize: messagesBufferSize,
		sessionPool:        sessionPool,
		metricsReporters:   metricsReporters,
		serializer:         serializer,
	}
}

// CreateAgent returns a new agent
func (f *agentFactoryImpl) CreateAgent(conn net.Conn) Agent {
	return newAgent(conn, f.packetCodec, f.serializer, f.heartbeatTimeout, f.writeTimeout, f.messagesBufferSize, f.appDieChan, f.messageCodec, f.metricsReporters, f.sessionPool)
}

// NewAgent create new agent instance
func newAgent(
	conn net.Conn,
	packetCodec codec.PacketCodec,
	serializer serialize.Serializer,
	heartbeatTime time.Duration,
	writeTimeout time.Duration,
	messagesBufferSize int,
	dieChan chan bool,
	messageCodec message.MessageCodec,
	metricsReporters []metrics.Reporter,
	sessionPool session.SessionPool,
) Agent {
	// initialize heartbeat and handshake data on first user connection
	serializerName := serializer.GetName()

	once.Do(func() {
		hbdEncode(heartbeatTime, packetCodec, messageCodec.IsCompressionEnabled(), serializerName)
		herdEncode(heartbeatTime, packetCodec, messageCodec.IsCompressionEnabled(), serializerName)
	})

	if writeTimeout <= 0 {
		writeTimeout = config.DefaultWriteTimeout
	}

	a := &agentImpl{
		appDieChan:         dieChan,
		chDie:              make(chan struct{}),
		chSend:             make(chan pendingWrite, messagesBufferSize),
		chStopHeartbeat:    make(chan struct{}),
		chStopWrite:        make(chan struct{}),
		messagesBufferSize: messagesBufferSize,
		conn:               conn,
		packetCodec:        packetCodec,
		heartbeatTimeout:   heartbeatTime,
		writeTimeout:       writeTimeout,
		lastAt:             time.Now().Unix(),
		serializer:         serializer,
		state:              constants.StatusStart,
		messageCodec:       messageCodec,
		metricsReporters:   metricsReporters,
		sessionPool:        sessionPool,
	}

	// binding session
	s := sessionPool.NewSession(a, true)
	metrics.ReportNumberOfConnectedClients(metricsReporters, sessionPool.GetSessionCount())
	a.Session = s
	return a
}

func (a *agentImpl) getMessageFromPendingMessage(pm pendingMessage) (*message.Message, error) {
	payload, err := util.SerializeOrRaw(a.serializer, pm.payload)
	if err != nil {
		payload, err = util.GetErrorPayload(a.serializer, err)
		if err != nil {
			return nil, err
		}
	}

	// construct message and encode
	m := &message.Message{
		Type:  pm.typ,
		Data:  payload,
		Route: pm.route,
		ID:    pm.mid,
		Err:   pm.err,
	}

	return m, nil
}

func (a *agentImpl) packetEncodeMessage(m *message.Message) ([]byte, error) {
	em, err := a.messageCodec.Encode(m)
	if err != nil {
		return nil, err
	}

	// packet encode
	p, err := a.packetCodec.Encode(packet.Data, em)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (a *agentImpl) send(pendingMsg pendingMessage) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.NewError(
				fmt.Errorf("%s: %s", constants.ErrBrokenPipe.Error(), panicErr),
				errors.ErrClientClosedRequest,
			)
			logger.Log.Error("agent send panicked: ", err)
		}
	}()
	a.reportChannelSize()

	m, err := a.getMessageFromPendingMessage(pendingMsg)
	if err != nil {
		logger.Log.Errorf(
			"agent send failed when getting pending msg. route: %s, type: %s, err: %s",
			pendingMsg.route, &pendingMsg.typ, err,
		)
		return err
	}

	// packet encode
	p, err := a.packetEncodeMessage(m)
	if err != nil {
		logger.Log.Errorf(
			"agent send failed when encoding the msg. route: %s, type: %s, err: %s",
			pendingMsg.route, &pendingMsg.typ, err,
		)
		return err
	}

	pWrite := pendingWrite{
		ctx:  pendingMsg.ctx,
		data: p,
	}

	if pendingMsg.err {
		pWrite.err = util.GetErrorFromPayload(a.serializer, m.Data)
	}

	// chSend is never closed so we need this to don't block if agent is already closed
	select {
	case a.chSend <- pWrite:
	case <-a.chDie:
	}
	return
}

// GetSession returns the agent session
func (a *agentImpl) GetSession() session.Session {
	return a.Session
}

// Push implementation for NetworkEntity interface
func (a *agentImpl) Push(route string, v interface{}) error {
	if a.GetStatus() == constants.StatusClosed {
		return errors.NewError(constants.ErrBrokenPipe, errors.ErrClientClosedRequest)
	}

	switch d := v.(type) {
	case []byte:
		logger.Log.Debugf("Type=Push, ID=%d, UID=%s, Route=%s, Data=%dbytes",
			a.Session.ID(), a.Session.UID(), route, len(d))
	default:
		logger.Log.Debugf("Type=Push, ID=%d, UID=%s, Route=%s, Data=%+v",
			a.Session.ID(), a.Session.UID(), route, v)
	}
	return a.send(pendingMessage{typ: message.Push, route: route, payload: v})
}

// ResponseMID implementation for NetworkEntity interface
// Respond message to session
func (a *agentImpl) ResponseMID(ctx context.Context, mid uint, v interface{}, isError ...bool) error {
	err := false
	if len(isError) > 0 {
		err = isError[0]
	}
	if a.GetStatus() == constants.StatusClosed {
		return errors.NewError(constants.ErrBrokenPipe, errors.ErrClientClosedRequest)
	}

	if mid <= 0 {
		return constants.ErrSessionOnNotify
	}

	switch d := v.(type) {
	case []byte:
		logger.Log.Debugf("Type=Response, ID=%d, UID=%s, MID=%d, Data=%dbytes",
			a.Session.ID(), a.Session.UID(), mid, len(d))
	default:
		logger.Log.Infof("Type=Response, ID=%d, UID=%s, MID=%d, Data=%+v",
			a.Session.ID(), a.Session.UID(), mid, v)
	}

	return a.send(pendingMessage{ctx: ctx, typ: message.Response, mid: mid, payload: v, err: err})
}

// Close closes the agent, cleans inner state and closes low-level connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (a *agentImpl) Close() error {
	a.closeMutex.Lock()
	defer a.closeMutex.Unlock()
	if a.GetStatus() == constants.StatusClosed {
		return constants.ErrCloseClosedSession
	}
	a.SetStatus(constants.StatusClosed)

	logger.Log.Debugf("Session closed, ID=%d, UID=%s, IP=%s",
		a.Session.ID(), a.Session.UID(), a.conn.RemoteAddr())

	// prevent closing closed channel
	select {
	case <-a.chDie:
		// expect
	default:
		close(a.chStopWrite)
		close(a.chStopHeartbeat)
		close(a.chDie)
		a.onSessionClosed(a.Session)
	}

	metrics.ReportNumberOfConnectedClients(a.metricsReporters, a.sessionPool.GetSessionCount())

	return a.conn.Close()
}

// RemoteAddr implementation for NetworkEntity interface
// returns the remote network address.
func (a *agentImpl) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

// String, implementation for Stringer interface
func (a *agentImpl) String() string {
	return fmt.Sprintf("Remote=%s, LastTime=%d", a.conn.RemoteAddr().String(), atomic.LoadInt64(&a.lastAt))
}

// GetStatus gets the status
func (a *agentImpl) GetStatus() int32 {
	return atomic.LoadInt32(&a.state)
}

// Kick sends a kick packet to a client
func (a *agentImpl) Kick(ctx context.Context) error {
	// packet encode
	p, err := a.packetCodec.Encode(packet.Kick, nil)
	if err != nil {
		return err
	}
	_, err = a.conn.Write(p)
	return err
}

// SetLastAt sets the last at to now
func (a *agentImpl) SetLastAt() {
	atomic.StoreInt64(&a.lastAt, time.Now().Unix())
}

// SetStatus sets the agent status
func (a *agentImpl) SetStatus(state int32) {
	atomic.StoreInt32(&a.state, state)
}

// Handle handles the messages from and to a client
func (a *agentImpl) Handle() {
	defer func() {
		a.Close()
		logger.Log.Debugf("Session handle goroutine exit, SessionID=%d, UID=%s", a.Session.ID(), a.Session.UID())
	}()

	go a.write()
	go a.heartbeat()
	<-a.chDie // agent closed signal
}

// IPVersion returns the remote address ip version.
// net.TCPAddr and net.UDPAddr implementations of String()
// always construct result as <ip>:<port> on both
// ipv4 and ipv6. Also, to see if the ip is ipv6 they both
// check if there is a colon on the string.
// So checking if there are more than one colon here is safe.
func (a *agentImpl) IPVersion() string {
	version := constants.IPv4

	ipPort := a.RemoteAddr().String()
	if strings.Count(ipPort, ":") > 1 {
		version = constants.IPv6
	}

	return version
}

func (a *agentImpl) heartbeat() {
	ticker := time.NewTicker(a.heartbeatTimeout)

	defer func() {
		ticker.Stop()
		a.Close()
	}()

	for {
		select {
		case <-ticker.C:
			deadline := time.Now().Add(-2 * a.heartbeatTimeout).Unix()
			if atomic.LoadInt64(&a.lastAt) < deadline {
				logger.Log.Debugf("Session heartbeat timeout, LastTime=%d, Deadline=%d", atomic.LoadInt64(&a.lastAt), deadline)
				return
			}

			// chSend is never closed so we need this to don't block if agent is already closed
			select {
			case a.chSend <- pendingWrite{data: hbd}:
			case <-a.chDie:
				return
			case <-a.chStopHeartbeat:
				return
			}
		case <-a.chDie:
			return
		case <-a.chStopHeartbeat:
			return
		}
	}
}

func (a *agentImpl) onSessionClosed(s session.Session) {
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Errorf("pitaya/onSessionClosed: %v", err)
		}
	}()

	for _, fn1 := range s.GetOnCloseCallbacks() {
		fn1()
	}

	for _, fn2 := range a.sessionPool.GetSessionCloseCallbacks() {
		fn2(s)
	}
}

// SendHandshakeResponse sends a handshake response
func (a *agentImpl) SendHandshakeResponse() error {
	_, err := a.conn.Write(hrd)

	return err
}

func (a *agentImpl) SendHandshakeErrorResponse() error {
	_, err := a.conn.Write(herd)

	return err
}

func (a *agentImpl) write() {
	// clean func
	defer func() {
		a.Close()
	}()

	for {
		select {
		case pWrite := <-a.chSend:
			ctx, err, data := pWrite.ctx, pWrite.err, pWrite.data

			writeErr := a.writeToConnection(ctx, data)

			tracing.FinishSpan(ctx, nil)

			if writeErr != nil {
				if e.Is(writeErr, os.ErrDeadlineExceeded) {
					// Log the timeout error but continue processing
					logger.Log.Warnf(
						"Context deadline exceeded for write in conn (%s) | session (%s): %s",
						a.conn.RemoteAddr(), a.Session.UID(), writeErr.Error(),
					)
				} else {
					err = errors.NewError(writeErr, errors.ErrClosedRequest)
					logger.Log.Errorf(
						"Failed to write in conn (%s) | session (%s): %s, agent will close",
						a.conn.RemoteAddr(), a.Session.UID(), writeErr.Error(),
					)
					metrics.ReportTimingFromCtx(ctx, a.metricsReporters, handlerType, err)
					// close agent if low-level conn broke during write
					return
				}
			}

			metrics.ReportTimingFromCtx(ctx, a.metricsReporters, handlerType, err)
		case <-a.chStopWrite:
			return
		}
	}
}

func (a *agentImpl) writeToConnection(ctx context.Context, data []byte) error {
	span := createConnectionSpan(ctx, a.conn, "conn write")

	a.conn.SetWriteDeadline(time.Now().Add(a.writeTimeout))
	_, writeErr := a.conn.Write(data)

	if span != nil {
		defer span.End()

		if writeErr != nil {
			span.RecordError(writeErr)
			span.SetStatus(codes.Error, writeErr.Error())
		}
	}

	return writeErr
}

func createConnectionSpan(ctx context.Context, conn net.Conn, op string) trace.Span {
	if ctx == nil {
		_, span := noop.NewTracerProvider().Tracer("noop").Start(context.TODO(), op)
		return span
	}
	remoteAddress := ""
	if conn.RemoteAddr() != nil {
		remoteAddress = conn.RemoteAddr().String()
	}

	attrs := []attribute.KeyValue{
		attribute.String("span.kind", "connection"),
		attribute.String("addr", remoteAddress),
	}

	_, span := tracing.StartSpan(ctx, op, attrs...)
	return span
}

// SendRequest sends a request to a server
func (a *agentImpl) SendRequest(ctx context.Context, serverID, route string, v interface{}) (*protos.Response, error) {
	return nil, e.New("not implemented")
}

// AnswerWithError answers with an error
func (a *agentImpl) AnswerWithError(ctx context.Context, mid uint, err error) {
	var e error
	defer func() {
		if e != nil {
			tracing.FinishSpan(ctx, e)
			metrics.ReportTimingFromCtx(ctx, a.metricsReporters, handlerType, e)
		}
	}()
	if ctx != nil && err != nil {
		span := trace.SpanFromContext(ctx)
		if span.IsRecording() {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
		}
	}
	p, e := util.GetErrorPayload(a.serializer, err)
	if e != nil {
		logger.Log.Errorf("error answering the user with an error: %s", e.Error())
		return
	}
	e = a.Session.ResponseMID(ctx, mid, p, true)
	if e != nil {
		logger.Log.Errorf("error answering the user with an error: %s", e.Error())
	}
}

func hbdEncode(heartbeatTimeout time.Duration, packetCodec codec.PacketCodec, dataCompression bool, serializerName string) {
	hData := map[string]interface{}{
		"code": 200,
		"sys": map[string]interface{}{
			"heartbeat":  heartbeatTimeout.Seconds(),
			"dict":       message.GetDictionary(),
			"serializer": serializerName,
		},
	}

	data, err := encodeAndCompress(hData, dataCompression)
	if err != nil {
		panic(err)
	}

	hrd, err = packetCodec.Encode(packet.Handshake, data)
	if err != nil {
		panic(err)
	}

	hbd, err = packetCodec.Encode(packet.Heartbeat, nil)
	if err != nil {
		panic(err)
	}
}

func herdEncode(heartbeatTimeout time.Duration, packetCodec codec.PacketCodec, dataCompression bool, serializerName string) {
	hErrData := map[string]interface{}{
		"code": 400,
		"sys": map[string]interface{}{
			"heartbeat":  heartbeatTimeout.Seconds(),
			"dict":       message.GetDictionary(),
			"serializer": serializerName,
		},
	}

	errData, err := encodeAndCompress(hErrData, dataCompression)
	if err != nil {
		panic(err)
	}

	herd, err = packetCodec.Encode(packet.Handshake, errData)
	if err != nil {
		panic(err)
	}
}

func encodeAndCompress(data interface{}, dataCompression bool) ([]byte, error) {
	encData, err := gojson.Marshal(data)
	if err != nil {
		return nil, err
	}

	if dataCompression {
		compressedData, err := compression.DeflateData(encData)
		if err != nil {
			return nil, err
		}

		if len(compressedData) < len(encData) {
			encData = compressedData
		}
	}
	return encData, nil
}

func (a *agentImpl) reportChannelSize() {
	chSendCapacity := a.messagesBufferSize - len(a.chSend)
	if chSendCapacity == 0 {
		logger.Log.Warnf("chSend is at maximum capacity")
	}
	for _, mr := range a.metricsReporters {
		if err := mr.ReportGauge(metrics.ChannelCapacity, map[string]string{"channel": "agent_chsend"}, float64(chSendCapacity)); err != nil {
			logger.Log.Warnf("failed to report chSend channel capaacity: %s", err.Error())
		}
	}
}

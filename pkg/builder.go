package pitaya

import (
	"github.com/google/uuid"
	"ricebean/pkg/acceptor"
	"ricebean/pkg/agent"
	"ricebean/pkg/cluster"
	"ricebean/pkg/config"
	"ricebean/pkg/conn/codec"
	"ricebean/pkg/conn/message"
	"ricebean/pkg/defaultpipelines"
	"ricebean/pkg/groups"
	"ricebean/pkg/logger"
	"ricebean/pkg/metrics"
	"ricebean/pkg/metrics/models"
	"ricebean/pkg/pipeline"
	"ricebean/pkg/router"
	"ricebean/pkg/serialize"
	"ricebean/pkg/service"
	"ricebean/pkg/session"
	"ricebean/pkg/worker"
)

// Builder holds dependency instances for a pitaya App
type Builder struct {
	acceptors      []acceptor.Acceptor
	postBuildHooks []func(app Pitaya)
	Config         config.PitayaConfig
	DieChan        chan bool
	//数据包编解码器
	PacketCodec  codec.PacketCodec
	MessageCodec message.MessageCodec
	Serializer   serialize.Serializer

	Router           *router.Router
	RPCClient        cluster.RPCClient
	RPCServer        cluster.RPCServer
	MetricsReporters []metrics.Reporter
	Server           *cluster.Server
	ServerMode       ServerMode
	ServiceDiscovery cluster.ServiceDiscovery
	Groups           groups.GroupService
	SessionPool      session.SessionPool
	Worker           *worker.Worker
	RemoteHooks      *pipeline.RemoteHooks
	HandlerHooks     *pipeline.HandlerHooks
}

// PitayaBuilder Builder interface
type PitayaBuilder interface {
	// AddPostBuildHook adds a post-build hook to the builder, a function receiving a Pitaya instance as parameter.
	AddPostBuildHook(hook func(app Pitaya))
	Build() Pitaya
}

// NewBuilderWithConfigs return a builder instance with default dependency instances for a pitaya App
// with configs defined by a config file (config.Config) and default paths (see documentation).
func NewBuilderWithConfigs(
	isFrontend bool,
	serverType string,
	serverMode ServerMode,
	serverMetadata map[string]string,
	conf *config.Config,
) *Builder {
	pitayaConfig := config.NewPitayaConfig(conf)
	return NewBuilder(
		isFrontend,
		serverType,
		serverMode,
		serverMetadata,
		*pitayaConfig,
	)
}

// NewDefaultBuilder return a builder instance with default dependency instances for a pitaya App,
// with default configs
func NewDefaultBuilder(isFrontend bool, serverType string, serverMode ServerMode, serverMetadata map[string]string, pitayaConfig config.PitayaConfig) *Builder {
	return NewBuilder(
		isFrontend,
		serverType,
		serverMode,
		serverMetadata,
		pitayaConfig,
	)
}

// NewBuilder return a builder instance with default dependency instances for a pitaya App,
// with configs explicitly defined
func NewBuilder(isFrontend bool,
	serverType string,
	serverMode ServerMode,
	serverMetadata map[string]string,
	config config.PitayaConfig,
) *Builder {
	server := cluster.NewServer(uuid.New().String(), serverType, isFrontend, serverMetadata)
	dieChan := make(chan bool)

	metricsReporters := []metrics.Reporter{}
	if config.Metrics.Prometheus.Enabled {
		metricsReporters = addDefaultPrometheus(config.Metrics, config.Metrics.Custom, metricsReporters, serverType)
	}

	if config.Metrics.Statsd.Enabled {
		metricsReporters = addDefaultStatsd(config.Metrics, metricsReporters, serverType)
	}

	handlerHooks := pipeline.NewHandlerHooks()
	if config.DefaultPipelines.StructValidation.Enabled {
		configureDefaultPipelines(handlerHooks)
	}

	sessionPool := session.NewSessionPool()

	var serviceDiscovery cluster.ServiceDiscovery
	var rpcServer cluster.RPCServer
	var rpcClient cluster.RPCClient
	if serverMode == Cluster {
		var err error
		serviceDiscovery, err = cluster.NewEtcdServiceDiscovery(config.Cluster.SD.Etcd, server, dieChan)
		if err != nil {
			logger.Log.Fatalf("error creating default cluster service discovery component: %s", err.Error())
		}

		rpcServer, err = cluster.NewNatsRPCServer(config.Cluster.RPC.Server.Nats, server, metricsReporters, dieChan, sessionPool)
		if err != nil {
			logger.Log.Fatalf("error setting default cluster rpc server component: %s", err.Error())
		}

		rpcClient, err = cluster.NewNatsRPCClient(config.Cluster.RPC.Client.Nats, server, metricsReporters, dieChan)
		if err != nil {
			logger.Log.Fatalf("error setting default cluster rpc client component: %s", err.Error())
		}
	}

	worker, err := worker.NewWorker(config.Worker, config.Worker.Retry)
	if err != nil {
		logger.Log.Fatalf("error creating default worker: %s", err.Error())
	}

	gsi := groups.NewMemoryGroupService(config.Groups.Memory)
	if err != nil {
		panic(err)
	}

	serializer, err := serialize.NewSerializer(serialize.Type(config.SerializerType))
	if err != nil {
		logger.Log.Fatalf("error creating serializer: %s", err.Error())
	}

	return &Builder{
		acceptors:        []acceptor.Acceptor{},
		postBuildHooks:   make([]func(app Pitaya), 0),
		Config:           config,
		DieChan:          dieChan,
		MessageCodec:     message.NewPomeloPacketEncoder(config.Handler.Messages.Compression),
		Serializer:       serializer,
		Router:           router.New(),
		RPCClient:        rpcClient,
		RPCServer:        rpcServer,
		MetricsReporters: metricsReporters,
		Server:           server,
		ServerMode:       serverMode,
		Groups:           gsi,
		RemoteHooks:      pipeline.NewRemoteHooks(),
		HandlerHooks:     handlerHooks,
		ServiceDiscovery: serviceDiscovery,
		SessionPool:      sessionPool,
		Worker:           worker,
	}
}

// AddAcceptor adds a new acceptor to app
func (builder *Builder) AddAcceptor(ac acceptor.Acceptor) {
	if !builder.Server.Frontend {
		logger.Log.Error("tried to add an acceptor to a backend server, skipping")
		return
	}
	builder.acceptors = append(builder.acceptors, ac)
}

// AddPostBuildHook adds a post-build hook to the builder, a function receiving a Pitaya instance as parameter.
func (builder *Builder) AddPostBuildHook(hook func(app Pitaya)) {
	builder.postBuildHooks = append(builder.postBuildHooks, hook)
}

// Build returns a valid App instance
func (builder *Builder) Build() Pitaya {
	handlerPool := service.NewHandlerPool()
	var remoteService *service.RemoteService
	if builder.ServerMode == Standalone {
		if builder.ServiceDiscovery != nil || builder.RPCClient != nil || builder.RPCServer != nil {
			panic("Standalone mode can't have RPC or service discovery instances")
		}
	} else {
		if !(builder.ServiceDiscovery != nil && builder.RPCClient != nil && builder.RPCServer != nil) {
			panic("Cluster mode must have RPC and service discovery instances")
		}

		builder.Router.SetServiceDiscovery(builder.ServiceDiscovery)
		//远程服务
		remoteService = service.NewRemoteService(
			builder.RPCClient,
			builder.RPCServer,
			builder.ServiceDiscovery,
			builder.PacketCodec,
			builder.Serializer,
			builder.Router,
			builder.MessageCodec,
			builder.Server,
			builder.SessionPool,
			builder.RemoteHooks,
			builder.HandlerHooks,
			handlerPool,
		)

		builder.RPCServer.SetPitayaServer(remoteService)
	}
	//代理工厂
	agentFactory := agent.NewAgentFactory(builder.DieChan,
		builder.PacketCodec,
		builder.Serializer,
		builder.Config.Heartbeat.Interval,
		builder.Config.Buffer.Agent.WriteTimeout,
		builder.MessageCodec,
		builder.Config.Buffer.Agent.Messages,
		builder.SessionPool,
		builder.MetricsReporters,
	)
	//接收消息服务器
	handlerService := service.NewHandlerService(
		builder.PacketCodec, //数据包解码器
		builder.MessageCodec,
		builder.Serializer,
		builder.Config.Buffer.Handler.LocalProcess,
		builder.Config.Buffer.Handler.RemoteProcess,
		builder.Server,
		remoteService,
		agentFactory,
		builder.MetricsReporters,
		builder.HandlerHooks,
		handlerPool,
	)

	app := NewApp(
		builder.ServerMode,
		builder.Serializer,
		builder.acceptors,
		builder.DieChan,
		builder.Router,
		builder.Server,
		builder.RPCClient,
		builder.RPCServer,
		builder.Worker,
		builder.ServiceDiscovery,
		remoteService,
		handlerService,
		builder.Groups,
		builder.SessionPool,
		builder.MetricsReporters,
		builder.Config,
	)

	for _, postBuildHook := range builder.postBuildHooks {
		postBuildHook(app)
	}

	return app
}

// NewDefaultApp returns a default pitaya app instance
func NewDefaultApp(isFrontend bool, serverType string, serverMode ServerMode, serverMetadata map[string]string, config config.PitayaConfig) Pitaya {
	builder := NewDefaultBuilder(isFrontend, serverType, serverMode, serverMetadata, config)
	return builder.Build()
}

func configureDefaultPipelines(handlerHooks *pipeline.HandlerHooks) {
	handlerHooks.BeforeHandler.PushBack(defaultpipelines.StructValidatorInstance.Validate)
}

func addDefaultPrometheus(config config.MetricsConfig, customMetrics models.CustomMetricsSpec, reporters []metrics.Reporter, serverType string) []metrics.Reporter {
	prometheus, err := CreatePrometheusReporter(serverType, config, customMetrics)
	if err != nil {
		logger.Log.Errorf("failed to start prometheus metrics reporter, skipping %v", err)
	} else {
		reporters = append(reporters, prometheus)
	}
	return reporters
}

func addDefaultStatsd(config config.MetricsConfig, reporters []metrics.Reporter, serverType string) []metrics.Reporter {
	statsd, err := CreateStatsdReporter(serverType, config)
	if err != nil {
		logger.Log.Errorf("failed to start statsd metrics reporter, skipping %v", err)
	} else {
		reporters = append(reporters, statsd)
	}
	return reporters
}

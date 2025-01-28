package toyblc

import (
	"context"
	"time"

	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/onexstack/onexstack/pkg/server"

	"github.com/onexstack/onex/internal/toyblc/pkg/blc"
	"github.com/onexstack/onex/internal/toyblc/pkg/miner"
	"github.com/onexstack/onex/internal/toyblc/pkg/ws"
)

// Config contains application-related configurations.
type Config struct {
	Miner           bool
	MinMineInterval time.Duration
	Address         string
	Accounts        map[string]string
	HTTPOptions     *genericoptions.HTTPOptions
	TLSOptions      *genericoptions.TLSOptions
	P2PAddr         string
	Peers           []string
}

// Server represents the web server.
type Server struct {
	srv server.Server
}

// ServerConfig contains the core dependencies and configurations of the server.
type ServerConfig struct {
	cfg             *Config
	bs              *blc.BlockSet
	ss              *ws.Sockets
	p2paddr         string
	miner           bool
	minMineInterval time.Duration
	accounts        map[string]string
	peers           []string
}

// NewServer initializes and returns a new Server instance.
func (cfg *Config) NewServer(ctx context.Context) (*Server, error) {
	// Create the core server instance.
	srv, err := InitializeWebServer(cfg)
	if err != nil {
		return nil, err
	}

	return &Server{srv: srv}, nil
}

// Run starts the server and listens for termination signals.
// It gracefully shuts down the server upon receiving a termination signal.
func (s *Server) Run(ctx context.Context) error {
	return server.Serve(ctx, s.srv)
}

type AggregatorServer struct {
	serverConfig *ServerConfig
	blcsrv       server.Server
	p2psrv       server.Server
}

func NewAggregatorServer(serverConfig *ServerConfig) (server.Server, error) {
	blcsrv := serverConfig.NewBlockChainServer()
	p2psrv := serverConfig.NewP2PServer()
	return &AggregatorServer{blcsrv: blcsrv, p2psrv: p2psrv}, nil
}

func (s *AggregatorServer) RunOrDie() {
	if s.serverConfig.miner {
		miner.NewMiner(s.serverConfig.bs, s.serverConfig.ss, s.serverConfig.minMineInterval).Start()
	}

	go s.p2psrv.RunOrDie()
	go s.blcsrv.RunOrDie()
	ws.ConnectToPeers(context.Background(), s.serverConfig.bs, s.serverConfig.ss, s.serverConfig.peers)
}

func (s *AggregatorServer) GracefulStop(ctx context.Context) {
	s.blcsrv.GracefulStop(ctx)
	s.p2psrv.GracefulStop(ctx)
}

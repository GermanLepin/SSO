package grpc_app

import (
	"fmt"
	"log/slog"
	"net"
	authgrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type GRPCApp struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGRPCApp(
	log *slog.Logger,
	port int,
) *GRPCApp {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)

	return &GRPCApp{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (g *GRPCApp) MustRun() {
	if err := g.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (g *GRPCApp) Run() error {
	const op = "grpc_app.Run"

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	g.log.Info(
		"grpc server started",
		slog.String("addr", listener.Addr().String()),
	)

	if err := g.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops gRPC server.
func (g *GRPCApp) Stop() {
	const op = "grpc_app.Stop"

	g.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", g.port))

	g.gRPCServer.GracefulStop()
}

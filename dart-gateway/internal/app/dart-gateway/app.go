package gatewayapp

import (
	"context"

	pgcl "github.com/10Narratives/dart/pkg/components/databases/postgres"
	grpcsrv "github.com/10Narratives/dart/pkg/components/transport/grpc/server"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type App struct {
	cfg *Config
	log *zap.Logger

	grpcServer *grpcsrv.Component
	stateDB    *pgcl.Component
}

func NewApp(cfg *Config, log *zap.Logger) (*App, error) {
	return &App{
		cfg: cfg,
		log: log,
		grpcServer: grpcsrv.NewComponent(
			cfg.Transport.GRPCServer.Address,
			grpcsrv.WithServerOptions(
				grpc.ChainUnaryInterceptor(
					validator.UnaryServerInterceptor(),
					// logging.UnaryServerInterceptor(),
					recovery.UnaryServerInterceptor(),
				),
			),
		),
		// stateDB:    pgcl.NewComponent(cfg.Databases.StateDB.DSN),
	}, nil
}

func (a *App) Startup(ctx context.Context) error {
	errGroup, errCtx := errgroup.WithContext(ctx)

	a.log.Info("application startup started")
	defer a.log.Info("application startup ended")

	errGroup.Go(func() error { return a.grpcServer.Startup(errCtx) })
	// errGroup.Go(func() error { return a.stateDB.Startup(errCtx) })

	return errGroup.Wait()
}

func (a *App) Shutdown(ctx context.Context) error {
	errGroup, errCtx := errgroup.WithContext(ctx)

	a.log.Info("application shutdown started")
	defer a.log.Info("application shutdown ended")

	errGroup.Go(func() error { return a.grpcServer.Shutdown(errCtx) })
	// errGroup.Go(func() error { return a.grpcServer.Shutdown(errCtx) })

	return errGroup.Wait()
}

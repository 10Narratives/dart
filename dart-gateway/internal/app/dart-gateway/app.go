package gatewayapp

import (
	"context"

	projectrepo "github.com/10Narratives/dart/dart-gateway/internal/repositories/project"
	projectsrv "github.com/10Narratives/dart/dart-gateway/internal/services/project"
	pgcl "github.com/10Narratives/dart/pkg/components/databases/postgres"
	grpcsrv "github.com/10Narratives/dart/pkg/components/transport/grpc/server"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type App struct {
	cfg *Config
	log *zap.Logger

	projectRepo    *projectrepo.Repository
	projectService *projectsrv.Service

	grpcServer *grpcsrv.Component
	stateDB    *pgcl.Component
}

func NewApp(cfg *Config, log *zap.Logger) (*App, error) {
	stateDB := pgcl.NewComponent(cfg.Databases.StateDB.DSN)

	// projectRepo := &projectrepo.Repository{}

	// projectService := &projectsrv.Service{}

	return &App{
		cfg: cfg,
		log: log,
		// projectRepo:    projectRepo,
		// projectService: projectService,
		stateDB: stateDB,
		grpcServer: grpcsrv.NewComponent(
			cfg.Transport.GRPCServer.Address,
			// grpcsrv.WithServerOptions(
			// 	grpc.ChainUnaryInterceptor(
			// 		validator.UnaryServerInterceptor(),
			// 		recovery.UnaryServerInterceptor(),
			// 	),
			// ),
			// grpcsrv.WithServiceRegistrations(
			// 	projectapi.Registration(projectService),
			// ),
		),
	}, nil
}

func (a *App) Startup(ctx context.Context) error {
	errGroup, errCtx := errgroup.WithContext(ctx)

	a.log.Info("application startup started")
	defer a.log.Info("application startup ended")

	errGroup.Go(func() error { return a.grpcServer.Startup(errCtx) })
	errGroup.Go(func() error { return a.stateDB.Startup(errCtx) })

	return errGroup.Wait()
}

func (a *App) Shutdown(ctx context.Context) error {
	errGroup, errCtx := errgroup.WithContext(ctx)

	a.log.Info("application shutdown started")
	defer a.log.Info("application shutdown ended")

	errGroup.Go(func() error { return a.grpcServer.Shutdown(errCtx) })
	errGroup.Go(func() error { return a.grpcServer.Shutdown(errCtx) })

	return errGroup.Wait()
}

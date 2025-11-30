package schedulerapp

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type App struct {
	cfg *Config
	log *zap.Logger
}

func NewApp(cfg *Config, log *zap.Logger) (*App, error) {
	return &App{
		cfg: cfg,
		log: log,
	}, nil
}

func (a *App) Startup(ctx context.Context) error {
	a.log.Debug("working...")
	time.Sleep(time.Minute)
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	return nil
}

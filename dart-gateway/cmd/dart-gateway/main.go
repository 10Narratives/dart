package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	gatewayapp "github.com/10Narratives/dart/dart-gateway/internal/app/dart-gateway"
	"github.com/10Narratives/dart/pkg/config"
	"go.uber.org/zap"
)

func main() {
	cfg, err := newConfig()
	if err != nil {
		fmt.Printf("cannot read configuration: %v\n", err)
		os.Exit(1)
	}

	log, err := newLogger(cfg.Environment)
	if err != nil {
		fmt.Printf("cannot initialize logger: %v", err)
		os.Exit(1)
	}

	log.Debug("initializing dart-gateway application")
	application, err := gatewayapp.NewApp(cfg, log)
	if err != nil {
		fmt.Printf("cannot initialize application: %v", err)
		os.Exit(1)
	}

	signalContext, signalCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer signalCancel()

	log.Debug("running dart-gateway application")

	go func() {
		if err := application.Startup(signalContext); err != nil {
			log.Error("fatal error during application running", zap.Error(err))
			signalCancel()
			return
		}
	}()

	<-signalContext.Done()
	log.Debug("shutdown signal received")

	shutdownContext, shutdownCancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer shutdownCancel()

	log.Debug("stopping dart-gateway application")
	if err := application.Shutdown(shutdownContext); err != nil {
		log.Error("fatal error during application stopping", zap.Error(err))
		os.Exit(1)
	}

	log.Debug("application stopped gracefully")
	os.Exit(0)
}

func newConfig() (*gatewayapp.Config, error) {
	var path string

	flag.StringVar(&path, "config", "", "path to configuration file")
	flag.Parse()

	if path == "" {
		return nil, errors.New("empty path to configuration file")
	}

	return config.ReadFromFile[gatewayapp.Config](path)
}

func newLogger(env string) (*zap.Logger, error) {
	switch env {
	case "prod":
		return zap.NewProduction()
	case "dev":
		return zap.NewDevelopment()
	default:
		return nil, errors.New("unsupported environment")
	}
}

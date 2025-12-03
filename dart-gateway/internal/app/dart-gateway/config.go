package gatewayapp

import (
	"time"
)

type Config struct {
	Environment     string          `yaml:"environment" env-default:"dev"`
	ShutdownTimeout time.Duration   `yaml:"shutdown_timeout" env-default:"5s"`
	Transport       TransportConfig `yaml:"transport"`
	Databases       DatabasesConfig `yaml:"databases"`
}

type TransportConfig struct {
	GRPCServer GRPCServerConfig `yaml:"grpc"`
}

type GRPCServerConfig struct {
	Address string `yaml:"address" env-required:"true"`
}

type DatabasesConfig struct {
	StateDB StateDBConfig `yaml:"statedb"`
}

type StateDBConfig struct {
	DSN string `yaml:"dsn"`
}

package schedulerapp

import "time"

type Config struct {
	Environment     string        `yaml:"environment" env-default:"dev"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env-default:"5s"`
}

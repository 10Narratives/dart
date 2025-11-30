package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func ReadFromFile[T any](path string) (*T, error) {
	stat, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, errors.New("file not found")
		}

		return nil, err
	}

	if stat.IsDir() {
		return nil, errors.New("expected regular file")
	}

	var cfg T
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("cannot parse configuration: %w", err)
	}

	return &cfg, nil
}

package configs

import (
	"fmt"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
)

type DatabaseConnConfig struct {
	Address string `yaml:"address" env-default:"localhost"`
	Port uint `yaml:"port" env-default:"5432"`
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password"`
	DBName string `yaml:"db_name" env-required:"true"`
	ConnTimeout time.Duration `yaml:"conn_timeout"`
}

type AuthConfig struct {
	Secret string `yaml:"secret" env-required:"true"`
	TokenLifetime time.Duration `yaml:"token_lifetime" env-default:"5m"`
}

type HTTPServerConfig struct {
	Port uint `yaml:"port" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Config struct {
	DatabaseConnConfig `yaml:"database_conn"`
	AuthConfig `yaml:"auth_service"`
	HTTPServerConfig `yaml:"http_server"`
}

func Load(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, ErrNoConfigPath
        }

        if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, ErrIsNotExist
        }

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("%v: %w", ErrReadFail, err)
	}

	return &cfg, nil
}

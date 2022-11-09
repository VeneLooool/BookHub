package config

import (
	"errors"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Server    ServerConfig
	Postgres  PostgresConfig
	Logger    Logger
	Memcached Memcached
}

type ServerConfig struct {
	AppVersion   string
	Port         string
	Mode         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Debug        bool
}

type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

type Memcached struct {
	URL        string
	Expiration int32
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

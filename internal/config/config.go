package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GRPCServer GRPCServerConfig `mapstructure:"grpc_server"`
	Database   DatabaseConfig   `mapstructure:"database"`
}

type GRPCServerConfig struct {
	Addr string `mapstructure:"addr"`
}

type DatabaseConfig struct {
	URL           string `mapstructure:"url"`
	MigrationPath string `mapstructure:"migration_path"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

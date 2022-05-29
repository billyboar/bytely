package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port           int    `mapstructure:"PORT"`
	GRPCServerAddr string `mapstructure:"GRPC_SERVER_ADDR"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	Domain         string `mapstructure:"DOMAIN"`
}

func LoadConfig() (*Config, error) {
	var config Config
	viper.SetEnvPrefix("bytely")
	viper.AutomaticEnv()

	viper.BindEnv("PORT")
	viper.BindEnv("GRPC_SERVER_ADDR")
	viper.BindEnv("LOG_LEVEL")
	viper.BindEnv("DOMAIN")

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

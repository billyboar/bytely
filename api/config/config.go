package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                  int    `mapstructure:"PORT"`
	DatabaseConnectionStr string `mapstructure:"DATABASE_CONNECTION_STR"`
	LogLevel              string `mapstructure:"LOG_LEVEL"`
}

func LoadConfig() (*Config, error) {
	var config Config
	viper.SetEnvPrefix("bytely")
	viper.AutomaticEnv()

	viper.BindEnv("PORT")
	viper.BindEnv("DATABASE_CONNECTION_STR")
	viper.BindEnv("LOG_LEVEL")

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

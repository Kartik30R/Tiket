package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
}

func NewEnvConfig() (*EnvConfig, error) {
	viper.SetConfigFile(".env")   
	viper.AutomaticEnv()         

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg EnvConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

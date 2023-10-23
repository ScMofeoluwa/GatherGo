package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	RefreshSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)

	return
}

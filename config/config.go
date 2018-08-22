package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Config *viper.Viper
}

func New() *Config {
	v := viper.New()
	v.SetConfigName("env")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	return &Config{v}
}

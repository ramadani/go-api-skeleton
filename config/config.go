package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config contains library for configuration.
type Config struct {
	Config *viper.Viper
}

// Init the configuration using viper and returns viper.
func Init() *Config {
	vp := viper.New()
	vp.SetConfigName("env")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	return &Config{vp}
}

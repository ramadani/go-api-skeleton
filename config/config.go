package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config contains library for configuration.
type Config struct {
	vp *viper.Viper
}

func (cog *Config) GetInt(key string) int {
	return cog.vp.GetInt(key)
}

func (cog *Config) GetString(key string) string {
	return cog.vp.GetString(key)
}

func (cog *Config) GetBool(key string) bool {
	return cog.vp.GetBool(key)
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

package bootstrap

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct{}

func (c Config) Boot() {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func InitConfig() *Config {
	return &Config{}
}

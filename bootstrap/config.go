package bootstrap

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigBoot struct{}

func (b ConfigBoot) Boot() {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func NewConfigBoot() *ConfigBoot {
	return &ConfigBoot{}
}

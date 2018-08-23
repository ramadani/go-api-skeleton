package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("env")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	return vp
}

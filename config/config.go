package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadProperties() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./etc/config")
	viper.SetConfigName("env_local.yml")
	error := viper.ReadInConfig()
	if error != nil {
		panic(fmt.Errorf("fatal error loading config file: %w", error))
	}
}

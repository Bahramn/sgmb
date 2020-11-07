package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	App    AppConfig
	Server ServerConfiguration
	Log    LoggerConfig
}

func Load() (Configuration, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	// Set default value
	viper.SetDefault("log.path", "./storage/logs/")

	var c Configuration

	err = viper.Unmarshal(&c)

	return c, err
}

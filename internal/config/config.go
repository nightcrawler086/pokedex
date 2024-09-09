package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
}

func NewConfig() *Config {
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8080")
	viper.SetConfigName("configs/local")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			panic(err)
		}
	}
	return &Config{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}
}

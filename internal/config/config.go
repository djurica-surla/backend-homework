package config

import (
	"log"

	"github.com/spf13/viper"
)

// A struct which holds the app configuration
type Config struct {
	Port string `mapstructure:"port"`
	DSN  string `mapstructure:"dsn"`
}

var AppConfig *Config

// Function which reads configuration from config.json
func LoadAppConfig() {
	log.Println("loading server configuration...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}

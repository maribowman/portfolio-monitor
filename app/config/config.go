package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Config = loadConfig()

type config struct {
	Server ServerConfig
	Signal SignalConfig
}

func loadConfig() config {
	configFile := "local"
	configPath := "./configs"
	if len(os.Args[1:]) > 0 {
		if contains([]string{"test", "local", "int", "prod"}, os.Args[1]) {
			configFile = os.Args[1]
		} else {
			configPath = "./../../configs"
		}
	}
	viper.SetConfigName(configFile)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper error while trying to read '%s' file, %s", configFile, err)
	}
	var config config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
	}
	return config
}

func contains(values []string, key string) bool {
	for _, iter := range values {
		if iter == key {
			return true
		}
	}
	return false
}

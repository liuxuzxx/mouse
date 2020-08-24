package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server Server
}

type Server struct {
	Host string
	Name string
}

var Conf Config

func init() {
	log.Println("Load config properties information from config.yml file!")
	viper.AddConfigPath("./rattrap/config/")
	viper.SetConfigName("config")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Update the config file")
	})
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Parse the config file is error!")
	}
	err = viper.Unmarshal(&Conf)
}

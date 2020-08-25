package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   Server
	RocketMQ RocketMQ
	Database Database
}

type Server struct {
	Host string
	Name string
}

type RocketMQ struct {
	NameServer []string
	GroupName  string
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
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
	log.Printf("查看数据:%v\n", Conf)
}

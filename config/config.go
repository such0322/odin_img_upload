package config

import (
	"log"

	"github.com/go-ini/ini"
)

var configPath = "config/config.ini"
var GlobalConf Config

func Init() {
	var err error
	GlobalConf, err = ReadConfig(configPath)
	if err != nil {
		log.Println(err)
	}
}

type Config struct {
	Addr     string `ini:"addr"`
	NewsPath string `ini:"news_path"`
	Logger   bool   `ini:"logger_enable"`
}

func ReadConfig(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path)
	if err != nil {
		log.Println("load config fail")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config)
	if err != nil {
		log.Println("mapto struct fail")
		return config, err
	}

	return config, nil
}

package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Host   string
	Port   string
	DbHost string
	DbPort string
	DbName string
}

func InitConfigs() *Config {
	configPath := flag.String("config", "", "Path the config file")
	flag.Parse()
	var configs Config
	_, err := toml.DecodeFile(*configPath, &configs)
	if err != nil {
		log.Printf("Ошибка декодирования файла конфигов %v", err)
	}
	return &configs
}

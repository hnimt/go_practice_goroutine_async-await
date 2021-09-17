package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Server   server
	Database database
}

type server struct {
	Port string
}

type database struct {
	Type string
	Host string
	Port int64
	Name string
	User string
	Pass string
}

func Config() TomlConfig {
	var config TomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}

	return config
}

package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	CONFIG struct {
		APP    `yaml:"app"`
		HTTP   `yaml:"http"`
		SQLITE `yaml:"sqlite"`
	}

	APP struct {
		Name    string `yaml:"name" env:"APP_NAME" env-default:"go-study"`
		Version string `yaml:"version" env:"APP_VERSION" env-default:"v1.0.0"`
	}

	HTTP struct {
		URL  string `yaml:"url" env:"HTTP_URL" env-default:"http://127.0.0.1"`
		Port string `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
	}

	SQLITE struct {
		MAX_IDLE_CONNS int `yaml:"max_idle_conns" env:"SQLITE_MAX_IDLE_CONNS" env-default:"10"`
		MAX_OPEN_CONNS int `yaml:"max_open_conns" env:"SQLITE_MAX_OPEN_CONNS" env-default:"100"`
	}
)

func NewConfig(cfgFile string) *CONFIG {
	cfg := &CONFIG{}
	if err := cleanenv.ReadConfig(cfgFile, cfg); err != nil {
		fmt.Println(err)
	}
	return cfg
}

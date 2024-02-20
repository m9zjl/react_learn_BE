package config

import (
	"github.com/go-ini/ini"
	"log"
)

type AppConfig struct {
}

type ServerConfig struct {
	RunMode string
}

var cfg *ini.File

func Init() {
	var err error
	cfg, err = ini.Load("conf/pkg.ini")
	if err != nil {
		log.Fatalf("failed to load pkg.ini:%v", err)
	}

	mapTo("pkg", &AppConfig{})
	mapTo("server", &ServerConfig{})
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("failed to map %v, err:%v", section, err)
	}

}

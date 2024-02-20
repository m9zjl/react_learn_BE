package config

import (
	"github.com/go-ini/ini"
	"log"
	"server/app/config/vos"
)

var cfg *ini.File

var AppConfig = &vos.AppConfig{}

var ServerConfig = &vos.ServerConfig{}

func Init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("failed to load app.ini:%v", err)
	}

	mapTo("app", AppConfig)
	mapTo("server", ServerConfig)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("failed to map %v, err:%v", section, err)
	}

}

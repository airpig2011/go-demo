package utils

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type mongo struct {
	Url      string
	DbName      string
}

var MongoSetting = &mongo{}

var cfg *ini.File

func Setup() {

	//var mode string =  gin.Mode()
	var err error

	mode := os.Getenv("ONE_MODE")
	if mode == "dev"{
		log.Print("mode dev")
		cfg, err = ini.Load("conf/dev.ini")
	}else if mode == "pro"{
		log.Print("mode pro")
		cfg, err = ini.Load("conf/pro.ini")
	}else if mode == "test"{
		log.Print("mode test")
		cfg, err = ini.Load("conf/test.ini")
	}else{
		log.Print("mode default, please set env ONE_MODE")
		cfg, err = ini.Load("conf/dev.ini")
	}

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("mongo", MongoSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
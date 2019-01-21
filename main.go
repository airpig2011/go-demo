package main

import (
	"fmt"
	"log"
	"net/http"
	"one/go-misc/router"
	"one/go-misc/utils"
)

func main() {

	router := router.InitRouter()

	utils.Setup()

	log.Print(utils.MongoSetting.Url)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", utils.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    utils.ServerSetting.ReadTimeout,
		WriteTimeout:   utils.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
package router

import (
	"github.com/gin-gonic/gin"
	"one/go-misc/modules/map/controller"
)

func InitRouter() *gin.Engine {

	router := gin.New()



	//router.Use(gin.Logger())
	//
	//router.Use(gin.Recovery())
	//gin.SetMode(setting.ServerSetting.RunMode)

	router.GET("/test", controller.Test)

	return router
}


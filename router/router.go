package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/go-gin-template/controller"
)

func Route() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", controller.PingGet)
	r.POST("/ping", controller.PingPost)
	return r
}

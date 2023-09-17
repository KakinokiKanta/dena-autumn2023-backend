package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/controller"
)

func Route() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", controller.PingGet)
	r.POST("/ping", controller.PingPost)

	r.GET("/users/:user_id", controller.UserGet)
	r.GET("/users", controller.UsersGet)

	r.POST("/signup", controller.Signup)

	r.GET("/answer/:answer_id", controller.AnswerGet)
	r.GET("/answers", controller.AnswersGet)
	r.POST("/answer", controller.AnswerPost)

	r.GET("/theme/:theme_id", controller.ThemeGet)
	r.POST("/theme", controller.ThemePost)
	return r
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/go-gin-template/model"
)

func PingGet(c *gin.Context) {
	res := model.Ping{
		Message: "pong",
	}
	c.JSON(200, res)
}

func PingPost(c *gin.Context) {
	req := model.Ping{}

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err)
		return
	}

	res := model.Ping{
		Message: req.Message,
	}
	c.JSON(200, res)
}

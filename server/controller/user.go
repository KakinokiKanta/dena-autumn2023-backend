package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/model"
)

func UsersGet(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		fmt.Errorf("Cannot get user id")
	}
}

func UserPost(c *gin.Context) {
	userId := c.PostForm("user_id")
	name := c.PostForm("name")
}

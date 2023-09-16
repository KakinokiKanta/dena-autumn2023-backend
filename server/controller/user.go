package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/model"
)

func UserGet(c *gin.Context) {
	db := model.GetDB()
	userId := c.Param("user_id")

	user, err:= model.GetUserByUserID(db, userId)

	if err != nil {
		fmt.Println("Cannot get User")
	}
	c.JSON(200, user)
}

func UsersGet(c *gin.Context) {
	db := model.GetDB()

	users, err:= model.GetUsers(db)

	if err != nil {
		fmt.Println("Cannot get Users")
	}
	c.JSON(200, users)
}

func UserPost(c *gin.Context) {
	// userId := c.PostForm("user_id")
	// name := c.PostForm("name")
}

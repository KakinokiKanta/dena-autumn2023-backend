package controller

import (
	"fmt"
	"strconv"

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
		c.Error(err)
		return
	}
	c.JSON(200, users)
}

func UserPost(c *gin.Context) {
	db := model.GetDB()
	userId := c.PostForm("user_id")
	name := c.PostForm("name")
	isNew, err := strconv.ParseBool(c.PostForm("is_new"))
	if err != nil {
		c.Error(err)
		return
	}
	user := model.User{ID: userId, Name: name, IsNew: isNew}
	err = model.PutUser(db, user)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, "User Created")
}

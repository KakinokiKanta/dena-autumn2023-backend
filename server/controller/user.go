package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/model"
)

func UserGet(c *gin.Context) {
	db := model.GetDB()
	userId := c.Param("user_id")

	user, err:= model.GetUserByUserID(db, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer db.Close()
	}
	c.JSON(200, user)
	defer db.Close()
}

func UsersGet(c *gin.Context) {
	db := model.GetDB()

	users, err:= model.GetUsers(db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer db.Close()
		return
	}
	c.JSON(200, users)
	defer db.Close()
}

func UserPost(c *gin.Context) {
	db := model.GetDB()
	var jsonUser model.User
	if err := c.ShouldBindJSON(&jsonUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer db.Close()
		return
	}

	err := model.PutUser(db, jsonUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer db.Close()
		return
	}
	c.JSON(201, "User Created")
	defer db.Close()
}

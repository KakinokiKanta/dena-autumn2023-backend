package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/model"
)

func ThemeGet(c *gin.Context) {
	db := model.GetDB()
	themeId := c.Param("theme_id")
	theme, err := model.GetThemeByID(db, themeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer db.Close()
		return
	}
	c.JSON(200, theme)
	defer db.Close()
}

func ThemePost(c *gin.Context) {
	db := model.GetDB()
	var jsonTheme model.Theme
	if err := c.ShouldBindJSON(&jsonTheme); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer db.Close()
		return
	}

	err := model.PutTheme(db, jsonTheme)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer db.Close()
		return
	}
	c.JSON(201, "Theme Created")
	defer db.Close()
}

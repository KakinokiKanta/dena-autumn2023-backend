package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/model"
)

func ThemeGet(c *gin.Context) {
	db := model.GetDB()
	themeId := c.Param("theme_id")
	theme, err := model.GetThemeByID(db, themeId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, theme)
}

func ThemePost(c *gin.Context) {
	db := model.GetDB()
	themeId := c.PostForm("theme_id")
	name := c.PostForm("name")
	theme := model.Theme{ID: themeId, Name: name}
	err := model.PutTheme(db, theme)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, "Theme Created")
}

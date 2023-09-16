package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/shinbunbun/dena-autumn-backend/server/model"
)

func ThemeGet(c *gin.Context) {
	themeId, _ := strconv.Atoi(c.Param("theme_id"))
	fmt.Println(themeId)
}

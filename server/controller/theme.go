package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/shinbunbun/dena-autumn-backend/server/model"
)

func ThemeGet(c *gin.Context) {
	themeId, err := strconv.Atoi(c.Param("theme_id"))
	if err != nil {
		fmt.Errorf("Cannot get theme id")
		return
	}
	fmt.Println(themeId)
}

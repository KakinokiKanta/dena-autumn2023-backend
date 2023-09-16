package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/model"
)

func AnswerGet(c *gin.Context) {
	db := model.GetDB()
	answerId := c.Param("answer_id")
	answer, err := model.GetAnswerByID(db, answerId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, answer)
}

func AnswersGet(c *gin.Context) {
	db := model.GetDB()
	answers, err := model.GetAnswers(db)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, answers)
}

func AnswersGetByUserId(c *gin.Context) {
	db := model.GetDB()
	userId := c.Param("user_id")
	answers, err := model.GetAnswersByUserID(db, userId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, answers)
}

func AnswerPost(c *gin.Context) {
	db := model.GetDB()
	userId := c.PostForm("user_id")
	themeId := c.PostForm("theme_id")
	content := c.PostForm("content")
	answer := model.Answer{UserID: userId, ThemeID: themeId, Content: content}
	err := model.PutAnswer(db, answer)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, "Answer Created")
}

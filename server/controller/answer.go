package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinbunbun/dena-autumn-backend/server/model"
)

func AnswerGet(c *gin.Context) {
	db := model.GetDB()
	answerId := c.Param("answer_id")
	answer, err := model.GetAnswerByID(db, answerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, answer)
}

func AnswersGet(c *gin.Context) {
	db := model.GetDB()
	answers, err := model.GetAnswers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, answers)
}

func AnswersGetByUserId(c *gin.Context) {
	db := model.GetDB()
	userId := c.Param("user_id")
	answers, err := model.GetAnswersByUserID(db, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, answers)
}

func AnswerPost(c *gin.Context) {
	db := model.GetDB()
	var jsonAnswer model.Answer
	if err := c.ShouldBindJSON(&jsonAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := model.PutAnswer(db, jsonAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, "Answer Created")
}

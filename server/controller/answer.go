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
	// user_idがあるならuser_idで検索
	userId := c.Query("user_id")
	if userId != "" {
		answers, err := model.GetAnswersByUserID(db, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, answers)
		return
	}

	// user_nameがあるならuser_nameで検索
	userName := c.Query("user_name")
	if userName != "" {
		answers, err := model.GetAnswersByUserName(db, userName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, answers)
		return
	}

	// クエリパラメータがない場合は、全検索
	answers, err := model.GetAnswers(db)
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

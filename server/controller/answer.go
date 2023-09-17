package controller

import (
	"fmt"
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
		defer db.Close()
		return
	}
	c.JSON(200, answer)

	defer db.Close()
}

func AnswersGet(c *gin.Context) {
	db := model.GetDB()
	// user_idがあるならuser_idで検索
	userId := c.Query("user_id")
	if userId != "" {
		answers, err := model.GetAnswersByUserID(db, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			defer db.Close()
			return
		}
		c.JSON(200, answers)
		defer db.Close()
		return
	}

	// user_nameがあるならuser_nameで検索
	userName := c.Query("user_name")
	if userName != "" {
		answers, err := model.GetAnswersByUserName(db, userName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			defer db.Close()
			return
		}
		c.JSON(200, answers)
		defer db.Close()
		return
	}

	// クエリパラメータがない場合は、全検索
	answers, err := model.GetAnswers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer db.Close()
		return
	}
	c.JSON(200, answers)
	defer db.Close()
}

func AnswerPost(c *gin.Context) {
	db := model.GetDB()
	var jsonAnswerWithUserName model.AnswerWithUserName
	if err := c.ShouldBindJSON(&jsonAnswerWithUserName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer db.Close()
		return
	}

	// Post時にリクエストにuser_nameが含まれている時
	userName := jsonAnswerWithUserName.UserName
	if userName != "" {
		err := model.PutAnswerByUserName(db, jsonAnswerWithUserName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			defer db.Close()
			return
		}
		c.JSON(201, "Answer Created")
		defer db.Close()
		return
	} else {
		jsonAnswer := model.Answer{
			ID: jsonAnswerWithUserName.ID, Content: jsonAnswerWithUserName.Content, UserID: jsonAnswerWithUserName.UserID, ThemeID: jsonAnswerWithUserName.ThemeID,
		}
		err := model.PutAnswer(db, jsonAnswer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			defer db.Close()
			return
		}
		c.JSON(201, "Answer Created")
		defer db.Close()
	}
}

package model

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	ID      string `gorm:"primary_key"`
	Answer  string
	UserID  string
	ThemeID string
}

func (Answer) GetAnswers(db *gorm.DB) []Answer {
	var answers []Answer
	db.Find(&answers)
	return answers
}

func (Answer) GetAnswersByUserID(db *gorm.DB, userID string) []Answer {
	var answers []Answer
	db.Where("user_id = ?", userID).Find(&answers)
	return answers
}

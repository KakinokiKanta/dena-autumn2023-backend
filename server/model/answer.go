package model

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	ID      string `gorm:"primary_key"`
	Content  string
	UserID  string
	ThemeID string
}

func GetAnswers(db *gorm.DB) (answers []Answer, err error) {
	err = db.Find(&answers).Error
	return
}

func GetAnswerByID(db *gorm.DB, id string) (answer Answer, err error) {
	err = db.Where("id = ?", id).Find(&answer).Error
	return
}

func GetAnswersByUserID(db *gorm.DB, userID string) (answers []Answer, err error) {
	err = db.Where("user_id = ?", userID).Find(&answers).Error
	return
}

func PutAnswer(db *gorm.DB, answer Answer) (err error) {
	err = db.Create(&answer).Error
	return
}

package model

import (
	"fmt"

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

func GetAnswersByUserName(db *gorm.DB, name string) (answers []Answer, err error) {
	user := User{}
	err = db.Where("name = ?", name).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("Cannot get user by user name")
	}
	err = db.Where("user_id = ?", user.ID).Find(&answers).Error
	return
}

func PutAnswer(db *gorm.DB, answer Answer) (err error) {
	err = db.Create(&answer).Error
	return
}

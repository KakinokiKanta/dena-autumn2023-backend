package model

import (
	"math/rand"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oklog/ulid"
)

type Answer struct {
	ID      string `json:"id" gorm:"primary_key"`
	Content  string `json:"content"`
	UserID  string `json:"user_id"`
	ThemeID string `json:"theme_id"`
}

type AnswerWithUserName struct {
	ID      string `json:"id" gorm:"primary_key"`
	Content  string `json:"content"`
	UserID  string `json:"user_id"`
	UserName string `json:"user_name"`
	ThemeID string `json:"theme_id"`
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
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	if answer.ID == "" {
		answer.ID = id.String()
	}
	err = db.Create(&answer).Error
	return
}

func PutAnswerByUserName(db *gorm.DB, answerWithUserName AnswerWithUserName) (err error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	if answerWithUserName.ID == "" {
		answerWithUserName.ID = id.String()
	}
	user, err := GetUserByUserName(db, answerWithUserName.UserName)
	if err != nil {
		return
	}
	answer := Answer{ID: answerWithUserName.ID, Content: answerWithUserName.Content, UserID: user.ID, ThemeID: answerWithUserName.ThemeID}
	err = db.Create(&answer).Error
	return
}

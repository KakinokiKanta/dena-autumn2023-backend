package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID    string `gorm:"primary_key"`
	Name  string
	IsNew bool
}

func GetUserByUserID(db *gorm.DB, userID string) (user User, err error) {
	err = db.Where("id = ?", userID).First(&user).Error
	return
}

func GetUsers(db *gorm.DB) (users []User, err error) {
	err = db.Find(&users).Error
	return 
}

func (user User) PutUser(db *gorm.DB) {
	db.Create(&user)
}

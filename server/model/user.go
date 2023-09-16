package model

import "github.com/jinzhu/gorm"

type User struct {
	ID    string `gorm:"primary_key"`
	Name  string
	IsNew bool
}

func (User) GetUserByUserID(db *gorm.DB, userID string) (user User, err error) {
	err = db.Where("id = ?", userID).First(&user).Error
	return
}

func (User) PutUser(db *gorm.DB, user User) {
	db.Create(&user)
}

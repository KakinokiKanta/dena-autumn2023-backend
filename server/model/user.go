package model

import (
	"math/rand"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/oklog/ulid"
)

type User struct {
	ID    string `gorm:"primary_key;type:string;default:uuid_generate_v4()"`
	Name  string
	IsNew bool `json:"is_new"`
}

func GetUserByUserID(db *gorm.DB, userID string) (user User, err error) {
	err = db.Where("id = ?", userID).First(&user).Error
	return
}

func GetUsers(db *gorm.DB) (users []User, err error) {
	err = db.Find(&users).Error
	return 
}

func Signup(db *gorm.DB, user User) (err error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	if user.ID == "" {
		user.ID = id.String()
	}
	err = db.Create(&user).Error
	return err
}

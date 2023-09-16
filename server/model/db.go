package model

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:pass@tcp(localhost:3306)/my_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}

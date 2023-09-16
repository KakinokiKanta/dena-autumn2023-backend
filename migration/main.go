package main

import (
	"github.com/jinzhu/gorm"
	"github.com/shinbunbun/dena-autumn-backend/server/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/my_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.Theme{},
	)
	db.AutoMigrate(&model.Answer{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("theme_id", "themes(id)", "CASCADE", "CASCADE")
}

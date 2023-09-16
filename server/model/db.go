package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *gorm.DB {
	var dns string
	if os.Getenv("ENV") == "prd" {
		var (
			user                   = os.Getenv("DB_USER")
			pass                   = os.Getenv("DB_PASS")
			instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
			dbName                 = os.Getenv("DB_NAME")
			socketDir              = "cloudsql"
		)
		dns = fmt.Sprintf("%s:%s@unix(%s/%s)/%s?parseTime=true&charset=utf8", user, pass, socketDir, instanceConnectionName, dbName)
	} else {
		dns = "user:pass@tcp(localhost:3306)/my_db?charset=utf8&parseTime=True&loc=Local"
	}
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	return db
}

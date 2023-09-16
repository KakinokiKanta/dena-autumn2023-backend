package model

import "github.com/jinzhu/gorm"

type Theme struct {
	ID   string `gorm:"primary_key"`
	Name string
}

func GetThemeByID(db *gorm.DB, themeID string) (theme Theme, err error) {
	err = db.Where("id = ?", themeID).First(&theme).Error
	return
}

func PutTheme(db *gorm.DB, theme Theme) (err error) {
	err = db.Create(&theme).Error
	return
}

package model

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oklog/ulid"
)

type Theme struct {
	ID   string `gorm:"primary_key"`
	Name string
}

func GetThemeByID(db *gorm.DB, themeID string) (theme Theme, err error) {
	err = db.Where("id = ?", themeID).First(&theme).Error
	return
}

func PutTheme(db *gorm.DB, theme Theme) (err error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	if theme.ID == "" {
		theme.ID = id.String()
	}
	err = db.Create(&theme).Error
	return
}

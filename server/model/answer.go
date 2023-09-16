package model

type Answer struct {
	ID      string `gorm:"primary_key"`
	Answer  string
	UserID  string
	ThemeID string
}

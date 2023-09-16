package model

type Theme struct {
	ID   string `gorm:"primary_key"`
	Name string
}

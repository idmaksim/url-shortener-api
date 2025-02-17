package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	ShortURL    string `gorm:"unique"`
	OriginalURL string
}

package models

import "time"

// URL represents URL model
// @Description URL model contains original and shortened URL
type URL struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ShortURL    string `json:"shortURL" gorm:"unique" index:"idx_short_url"`
	OriginalURL string `json:"originalURL"`
}

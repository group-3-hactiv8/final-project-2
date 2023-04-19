package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_ID   int    `json:"user_id"`
}

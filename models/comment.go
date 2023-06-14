package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId  uint    `json:"user_id"`
	PhotoId uint    `json:"photo_id"`
	Message  string `json:"message"`
	User User
	Photo Photo
}

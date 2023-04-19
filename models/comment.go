package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	User_ID  int    `json:"user_id"`
	Photo_ID int    `json:"photo_id"`
	Message  string `json:"message"`
}

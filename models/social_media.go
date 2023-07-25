package models

import (
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	ID             uint   `gorm:"primary_key" json:"id"`
	Name           string `json:"name" gorm:"not null" valid:"required~Your age is required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" valid:"required~Your age is required, url~Invalid url format"`
	UserId         uint   `json:"user_id"`
	User           User
}


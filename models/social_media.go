package models

import (
	"final-project-2/pkg/errs"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint   `gorm:"primary_key" json:"id"`
	Name           string `json:"name" gorm:"not null" valid:"required~Your age is required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" valid:"required~Your age is required, url~Invalid url format"`
	UserId         uint   `json:"user_id"`
	User           *User
}

func (sc *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(sc)

	if err != nil {
		return errs.NewUnprocessableEntity(err.Error())
	}
	return nil
}

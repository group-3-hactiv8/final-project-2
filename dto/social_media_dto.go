package dto

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type NewSocialMediaRequest struct {
	Name           string `json:"name" valid:"required~Your Name is required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Your social media url is required, url~Invalid url format"`
}

func (newSM *NewSocialMediaRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(newSM)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (newSM *NewSocialMediaRequest) NewSocialMediaRequestToModel() *models.SocialMedia {
	return &models.SocialMedia{
		Name:           newSM.Name,
		SocialMediaUrl: newSM.SocialMediaUrl,
	}
}

type NewSocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserOfSocialMediaResponse struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}
type SocialMediaResponse struct {
	ID             uint                      `json:"id"`
	Name           string                    `json:"name"`
	SocialMediaUrl string                    `json:"social_media_url"`
	UserId         uint                      `json:"user_id"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
	User           UserOfSocialMediaResponse `json:"user"`
}
type AllSocialMediasResponse struct {
	SocialMedias []SocialMediaResponse `json:"social_medias"`
}

type UpdateSocialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type DeleteSocialMediaResponse struct {
	Message string `json:"message"`
}

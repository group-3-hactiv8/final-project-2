package social_media_repository

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
)

type SocialMediaRepository interface {
	CreateSocialMedia(sm *models.SocialMedia) (*models.SocialMedia, errs.MessageErr)
	GetAllSocialMedias() (*[]models.SocialMedia, uint, errs.MessageErr)
	UpdateSocialMedia(sm *models.SocialMedia) (*models.SocialMedia, errs.MessageErr)
	DeleteSocialMedia(id uint, sm_id uint) errs.MessageErr
}

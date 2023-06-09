package photo_repository

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
)

type PhotoRepository interface {
	CreatePhoto(user *models.User, photo *models.Photo) (*models.Photo, errs.MessageErr)
	GetAllPhotos() ([]models.Photo, errs.MessageErr)
	GetPhotoByID(id uint) (*models.Photo, errs.MessageErr)
	UpdatePhoto(oldPhoto *models.Photo, newPhoto *models.Photo) (*models.Photo, errs.MessageErr)
	DeletePhoto(id uint) errs.MessageErr
}

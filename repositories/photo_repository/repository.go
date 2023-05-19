package photo_repository

import (
	"final-project-2/pkg/errs"
)

type PhotoRepository interface {
	CreatePhoto(user *entity.User, photo *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetAllPhotos() ([]entity.Photo, errs.MessageErr)
	GetPhotoByID(id uint) (*entity.Photo, errs.MessageErr)
	UpdatePhoto(oldPhoto *entity.Photo, newPhoto *entity.Photo) (*entity.Photo, errs.MessageErr)
	DeletePhoto(id uint) errs.MessageErr
}

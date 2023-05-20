package photo_pg

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/photo_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type photoPg struct {
	db *gorm.DB
}

func NewPhotoPG(db *gorm.DB) photo_repository.PhotoRepository {
	return &photoPg{db: db}
}

func (p *photoPg) CreatePhoto(user *models.User, photo *models.Photo) (*models.Photo, errs.MessageErr) {
	if err := p.db.Model(user).Association("Photos").Append(photo); err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to create new photo")
	}

	return photo, nil
}

func (p *photoPg) GetAllPhotos() ([]models.Photo, errs.MessageErr) {
	var photos []models.Photo
	if err := p.db.Find(&photos).Error; err != nil {
		return nil, errs.NewInternalServerError("Failed to get all photos")
	}

	return photos, nil
}

func (p *photoPg) GetPhotoByID(id uint) (*models.Photo, errs.MessageErr) {
	var photo models.Photo
	if err := p.db.First(&photo, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Photo with id %d is not found", id))
	}

	return &photo, nil
}

func (p *photoPg) UpdatePhoto(oldPhoto *models.Photo, newPhoto *models.Photo) (*models.Photo, errs.MessageErr) {
	if err := p.db.Model(oldPhoto).Updates(newPhoto).Error; err != nil {
		return nil, errs.NewInternalServerError(fmt.Sprintf("Failed to update photo with id %d", oldPhoto.ID))
	}

	return oldPhoto, nil
}

func (p *photoPg) DeletePhoto(id uint) errs.MessageErr {
	if err := p.db.Delete(&models.Photo{}, id).Error; err != nil {
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete photo with id %d", id))
	}

	return nil
}

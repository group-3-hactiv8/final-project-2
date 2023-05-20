package services

import (
	"final-project-2/dto"
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/photo_repository"
	"final-project-2/repositories/user_repository"
)

// PhotoService is an interface for performing CRUD operations on photos.
type PhotoService interface {
	// CreatePhoto creates a new photo and returns a response or an error.
	CreatePhoto(user *models.User, payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr)
	// GetAllPhotos retrieves all photos and their associated user information and returns them as a slice or an error.
	GetAllPhotos() ([]dto.GetAllPhotosResponse, errs.MessageErr)
	// UpdatePhoto updates the photo with the given ID and returns a response or an error.
	UpdatePhoto(id uint, payload *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr)
	// DeletePhoto deletes the photo with the given ID and returns a response or an error.
	DeletePhoto(id uint) (*dto.DeletePhotoResponse, errs.MessageErr)
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
	userRepo  user_repository.UserRepository
}

// NewPhotoService creates a new PhotoService object with the given photo and user repositories.
func NewPhotoService(photoRepo photo_repository.PhotoRepository, userRepo user_repository.UserRepository) PhotoService {
	return &photoService{photoRepo: photoRepo, userRepo: userRepo}
}

// CreatePhoto creates a new photo with the provided details and returns a response or an error.
func (p *photoService) CreatePhoto(user *models.User, payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr) {
	photo := payload.ToModels()

	createdPhoto, err := p.photoRepo.CreatePhoto(user, photo)
	if err != nil {
		return nil, err
	}

	response := &dto.CreatePhotoResponse{
		ID:        int(createdPhoto.ID),
		Title:     createdPhoto.Title,
		Caption:   createdPhoto.Caption,
		PhotoURL:  createdPhoto.PhotoURL,
		UserID:    createdPhoto.UserID,
		CreatedAt: createdPhoto.CreatedAt,
	}

	return response, nil
}

// GetAllPhotos retrieves all photos and their associated user information and returns them as a slice or an error.
func (p *photoService) GetAllPhotos() ([]dto.GetAllPhotosResponse, errs.MessageErr) {
	photos, err := p.photoRepo.GetAllPhotos()
	if err != nil {
		return nil, err
	}

	response := []dto.GetAllPhotosResponse{}
	for _, photo := range photos {
		user, err := p.photoRepo.GetPhotoByID(uint(photo.UserID))
		if err != nil {
			return nil, err
		}

		response = append(response, dto.GetAllPhotosResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    uint(photo.UserID),
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: dto.UserData{
				Email:    user.Email,
				Username: user.Username,
			},
		})
	}

	return response, nil
}

// UpdatePhoto updates the photo with the given ID and returns a response or an error.
func (p *photoService) UpdatePhoto(id uint, payload *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr) {
	oldPhoto, err := p.photoRepo.GetPhotoByID(id)
	if err != nil {
		return nil, err
	}
	newPhoto := payload.ToModels()

	updatedPhoto, err2 := p.photoRepo.UpdatePhoto(oldPhoto, newPhoto)
	if err2 != nil {
		return nil, err2
	}

	response := &dto.UpdatePhotoResponse{
		ID:        updatedPhoto.ID,
		Title:     updatedPhoto.Title,
		Caption:   updatedPhoto.Caption,
		PhotoURL:  updatedPhoto.PhotoURL,
		UserID:    uint(updatedPhoto.UserID),
		UpdatedAt: updatedPhoto.UpdatedAt,
	}

	return response, nil
}

// DeletePhoto deletes the photo with the given ID and returns a response or an error.
func (p *photoService) DeletePhoto(id uint) (*dto.DeletePhotoResponse, errs.MessageErr) {
	if err := p.photoRepo.DeletePhoto(id); err != nil {
		return nil, err
	}

	response := &dto.DeletePhotoResponse{
		Message: "Your photo has been successfully deleted",
	}

	return response, nil
}

// The services package contains all of the service layer code for working with photos.
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
	// CreatePhoto creates a new photo with the provided details and returns a response or an error.
	CreatePhoto(user *models.User, payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr)

	// GetAllPhotos retrieves all photos and their associated user information and returns them as a slice or an error.
	GetAllPhotos() ([]dto.GetAllPhotosResponse, errs.MessageErr)

	// UpdatePhoto updates the photo with the given ID and returns a response or an error.
	UpdatePhoto(id uint, payload *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr)

	// DeletePhoto deletes the photo with the given ID and returns a response or an error.
	DeletePhoto(id uint) (*dto.DeletePhotoResponse, errs.MessageErr)
}

// This struct implements the PhotoService interface.
type photoService struct {
	photoRepo photo_repository.PhotoRepository
	userRepo  user_repository.UserRepository
}

// NewPhotoService creates a new PhotoService object with the given photo and user repositories.
func NewPhotoService(photoRepo photo_repository.PhotoRepository, userRepo user_repository.UserRepository) PhotoService {
	// Return a new photoService object with the given repositories.
	return &photoService{photoRepo: photoRepo, userRepo: userRepo}
}

// CreatePhoto creates a new photo with the provided details and returns a response or an error.
func (p *photoService) CreatePhoto(user *models.User, payload *dto.CreatePhotoRequest) (*dto.CreatePhotoResponse, errs.MessageErr) {
	// Convert the DTO payload to a models.Photo object.
	photo := payload.ToModels()

	// Create the photo using the provided repository and user.
	createdPhoto, err := p.photoRepo.CreatePhoto(user, photo)
	if err != nil {
		return nil, err
	}

	// Convert the created photo to a DTO response.
	response := &dto.CreatePhotoResponse{
		ID:        int(createdPhoto.ID),
		Title:     createdPhoto.Title,
		Caption:   createdPhoto.Caption,
		PhotoURL:  createdPhoto.PhotoURL,
		UserID:    int(createdPhoto.UserId),
		CreatedAt: createdPhoto.CreatedAt,
	}

	// Return the response.
	return response, nil
}

// GetAllPhotos retrieves all photos and their associated user information and returns them as a slice or an error.
func (p *photoService) GetAllPhotos() ([]dto.GetAllPhotosResponse, errs.MessageErr) {
	// Retrieve all photos from the repository.
	photos, err := p.photoRepo.GetAllPhotos()
	if err != nil {
		return nil, err
	}

	// Create the response array.
	response := []dto.GetAllPhotosResponse{}
	for _, photo := range photos {
		// Get the associated user information for each photo.
		_, err := p.photoRepo.GetPhotoByID(uint(photo.UserId))
		if err != nil {
			return nil, err
		}

		// Convert the photo to a DTO response and append it to the response array.
		response = append(response, dto.GetAllPhotosResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    uint(photo.UserId),
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
		})
	}

	// Return the response array.
	return response, nil
}

// UpdatePhoto updates the photo with the given ID and returns a response or an error.
func (p *photoService) UpdatePhoto(id uint, payload *dto.UpdatePhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr) {
	// Retrieve the old photo to update.
	oldPhoto, err := p.photoRepo.GetPhotoByID(id)
	if err != nil {
		return nil, err
	}

	// Convert the DTO payload to a models.Photo object.
	newPhoto := payload.ToModels()

	// Update the photo using the provided repository.
	updatedPhoto, err2 := p.photoRepo.UpdatePhoto(oldPhoto, newPhoto)
	if err2 != nil {
		return nil, err2
	}

	// Convert the updated photo to a DTO response.
	response := &dto.UpdatePhotoResponse{
		ID:        updatedPhoto.ID,
		Title:     updatedPhoto.Title,
		Caption:   updatedPhoto.Caption,
		PhotoURL:  updatedPhoto.PhotoURL,
		UserID:    uint(updatedPhoto.UserId),
		UpdatedAt: updatedPhoto.UpdatedAt,
	}

	// Return the response.
	return response, nil
}

// DeletePhoto deletes the photo with the given ID and returns a response or an error.
func (p *photoService) DeletePhoto(id uint) (*dto.DeletePhotoResponse, errs.MessageErr) {
	// Delete the photo using the provided repository.
	if err := p.photoRepo.DeletePhoto(id); err != nil {
		return nil, err
	}

	// Create the response object.
	response := &dto.DeletePhotoResponse{
		Message: "Your photo has been successfully deleted",
	}

	// Return the response.
	return response, nil
}

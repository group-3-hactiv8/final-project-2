package dto

import (
	"final-project-2/models"
	"time"
)

// CreatePhotoRequest defines the request body for creating a photo.
type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required,url"`
}

// ToModels converts CreatePhotoRequest to models.Photo type.
func (p *CreatePhotoRequest) ToModels() *models.Photo {
	return &models.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: "",
		UserID:   0,
	}
}

// CreatePhotoResponse defines the response format for creating a photo.
type CreatePhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAllPhotosResponse defines the response format for getting all photos.
type GetAllPhotosResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserData  `json:"user"`
}

// UserData defines the user-related information used in various responses.
type UserData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

// UpdatePhotoRequest defines the request body for updating a photo.
type UpdatePhotoRequest CreatePhotoRequest

// ToModels converts UpdatePhotoRequest to models.Photo type.
func (p *UpdatePhotoRequest) ToModels() *models.Photo {
	return &models.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: "",
		UserID:   0,
	}
}

// UpdatePhotoResponse defines the response format for updating a photo.
type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DeletePhotoResponse defines the response format for deleting a photo.
type DeletePhotoResponse struct {
	Message string `json:"message"`
}

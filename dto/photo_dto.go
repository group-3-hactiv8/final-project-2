package dto

import (
	"final-project-2/models"
	"time"
)

// CreatePhotoRequest defines the request body for creating a photo.
type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`         // Title of the photo - mandatory field
	Caption  string `json:"caption"`                          // Caption of the photo - optional field
	PhotoURL string `json:"photo_url" binding:"required,url"` // Photo URL is a mandatory and must be valid URL
}

// ToModels converts CreatePhotoRequest to models.Photo type.
func (p *CreatePhotoRequest) ToModels() *models.Photo {
	return &models.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: p.PhotoURL,
	}
}

// CreatePhotoResponse defines the response format for creating a photo.
type CreatePhotoResponse struct {
	ID        int       `json:"id"`         // The unique identifier of the created photo
	Title     string    `json:"title"`      // Title of the photo
	Caption   string    `json:"caption"`    // Caption of the photo
	PhotoURL  string    `json:"photo_url"`  // URL of the photo
	UserID    int       `json:"user_id"`    // User ID of the creator of the photo
	CreatedAt time.Time `json:"created_at"` // Time at which photo was created
}

// GetAllPhotosResponse defines the response format for getting all photos.
type GetAllPhotosResponse struct {
	ID        uint      `json:"id"`         // The unique identifier of the photo
	Title     string    `json:"title"`      // Title of the photo
	Caption   string    `json:"caption"`    // Caption of the photo
	PhotoURL  string    `json:"photo_url"`  // URL of the photo
	UserID    uint      `json:"user_id"`    // User ID of the creator of the photo
	CreatedAt time.Time `json:"created_at"` // Time at which photo was created
	UpdatedAt time.Time `json:"updated_at"` // Time at which photo was updated
	User      UserData  `json:"user"`       // User related data
	Email     UserData  `json:"email"`      // Email address related to user data
}

// UserData defines the user-related information used in various responses.
type UserData struct {
	Email    string `json:"email"`    // User's email address
	Username string `json:"username"` // User's username
}

// UpdatePhotoRequest defines the request body for updating a photo.
type UpdatePhotoRequest CreatePhotoRequest

// ToModels converts UpdatePhotoRequest to models.Photo type.
func (p *UpdatePhotoRequest) ToModels() *models.Photo {
	return &models.Photo{
		Title:    p.Title,
		Caption:  p.Caption,
		PhotoURL: "",
		UserId:   0,
	}
}

// UpdatePhotoResponse defines the response format for updating a photo.
type UpdatePhotoResponse struct {
	ID        uint      `json:"id"`         // The unique identifier of the photo
	Title     string    `json:"title"`      // Title of the photo
	Caption   string    `json:"caption"`    // Caption of the photo
	PhotoURL  string    `json:"photo_url"`  // URL of the photo
	UserID    uint      `json:"user_id"`    // User ID of the creator of the photo
	UpdatedAt time.Time `json:"updated_at"` // Time at which photo was updated
}

// DeletePhotoResponse defines the response format for deleting a photo.
type DeletePhotoResponse struct {
	Message string `json:"message"` // Message indicating whether or not the photo was deleted
}

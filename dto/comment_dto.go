package dto

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

type NewCommentRequest struct {
	Message string `json:"message" binding:"required" valid:"required~Your Type is required"`
	PhotoId uint    `json:"photo_id" binding:"required" valid:"required~Your Type is required"`
}

func (c *NewCommentRequest) CommentRequestToModel() *models.Comment {
	return &models.Comment{
		Message: c.Message,
		PhotoId: c.PhotoId,
	}	
}

func (u *NewCommentRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

type NewCommentResponse struct {
	ID uint `json:"id"`
	Message string `json:"message"`
	PhotoId uint `json:"photo_id"`
	UserId uint `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}


type UserDataForComment struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}

type PhotoDataForComment struct {
	ID 		  uint 		`json:"id"`
	Title     string    `json:"title"`      
	Caption   string    `json:"caption"`   
	PhotoURL  string    `json:"photo_url"`  
	UserId    uint      `json:"user_id"`    
}

type GetAllCommentResponse struct {
	ID 			uint 				`json:"id"`
	Message 	string 				`json:"message"`
	PhotoId 	uint 				`json:"photo_id"`
	UserId 		uint 				`json:"user_id"`
	UpdatedAt 	time.Time 			`json:"updated_at"`
	CreatedAt 	time.Time 			`json:"created_at"`
	User 		UserDataForComment 	`json:"User"`
	Photo 		PhotoDataForComment `json:"Photo"`
}

func GetAllComment(comment models.Comment, photo models.Photo) GetAllCommentResponse {
	response := GetAllCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		UpdatedAt: comment.UpdatedAt,
		CreatedAt: comment.CreatedAt,
		User: UserDataForComment{
			ID:       comment.User.ID,
			Email:    comment.User.Email,
			Username: comment.User.Username,
		},
		Photo: PhotoDataForComment{
			ID:       comment.Photo.ID,
			Title:    comment.Photo.Title,
			Caption:  comment.Photo.Caption,
			PhotoURL: comment.Photo.PhotoURL,
			UserId:   comment.Photo.UserId,
		},
	}

	return response
}

type UpdateCommentRequest struct {
	Message string `json:"message" valid:"required~Your Type is required"`
}

func (c *UpdateCommentRequest) CommentUpdateRequestToModel() *models.Comment {
	return &models.Comment{
		Message: c.Message,
	}	
}

func (u *UpdateCommentRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

type UpdateCommentResponse struct {
	ID uint `json:"id"`
	Message string `json:"message"`
	PhotoId uint `json:"photo_id"`
	UserId uint `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
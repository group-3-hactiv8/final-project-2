package services

import (
	"final-project-2/dto"
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/comment_repository"
	"final-project-2/repositories/photo_repository"
	"final-project-2/repositories/user_repository"
)

type CommentService interface {
	CreateComment(user *models.User, payload *dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr)
	GetAllComment() ([]dto.GetAllCommentResponse, errs.MessageErr)
	GetCommentByUserId(userId uint) ([]dto.GetAllCommentResponse, errs.MessageErr)
	GetCommentByPhotoId(photoId uint) ([]dto.GetAllCommentResponse, errs.MessageErr)
	UpdateComment(id uint, comment *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr)
	DeleteComment(id uint) (*dto.DeleteCommentResponse, errs.MessageErr)
}

type commentService struct {
	commentRepo comment_repository.CommentRepository
	photoRepo   photo_repository.PhotoRepository
	userRepo    user_repository.UserRepository
}

func NewCommentService(
	commentRepo comment_repository.CommentRepository,
	photoRepo photo_repository.PhotoRepository,
	userRepo user_repository.UserRepository,
) CommentService {
	return &commentService{
		commentRepo: commentRepo, photoRepo: photoRepo, userRepo: userRepo}
}

func (c *commentService) CreateComment(user *models.User, payload *dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr) {
	comment := payload.CommentRequestToModel()

	_, errPhoto := c.photoRepo.GetPhotoByID(uint(comment.PhotoId))

	if errPhoto != nil {
		return nil, errPhoto
	}

	createComment, err := c.commentRepo.CreateComment(user, comment)

	if err != nil {
		return nil, err
	}

	response := &dto.NewCommentResponse{
		ID:        createComment.ID,
		Message:   createComment.Message,
		PhotoId:   createComment.PhotoId,
		UserId:    createComment.UserId,
		CreatedAt: createComment.CreatedAt,
	}
	return response, nil
}

func (c *commentService) GetCommentByUserId(userId uint) ([]dto.GetAllCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetCommentByUserId(userId)

	if err != nil {
		return nil, err
	}

	var response []dto.GetAllCommentResponse
	for _, comment := range comments {
		user, err := c.userRepo.GetUserByIDComment(comment.UserId)
		if err != nil {
			return nil, err
		}

		photo, err := c.photoRepo.GetPhotoByID(comment.PhotoId)
		if err != nil {
			return nil, err
		}

		commentResponse := convertToCommentResponse(comment, user, photo)
		response = append(response, commentResponse)
	}
	return response, nil
}

func (c *commentService) GetCommentByPhotoId(photoId uint) ([]dto.GetAllCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetCommentByPhotoId(photoId)
	if err != nil {
		return nil, err
	}

	var response []dto.GetAllCommentResponse
	for _, comment := range comments {
		user, err := c.userRepo.GetUserByIDComment(comment.UserId)
		if err != nil {
			return nil, err
		}

		photo, err := c.photoRepo.GetPhotoByID(comment.PhotoId)
		if err != nil {
			return nil, err
		}

		commentResponse := convertToCommentResponse(comment, user, photo)
		response = append(response, commentResponse)
	}
	return response, nil
}

func (c *commentService) GetAllComment() ([]dto.GetAllCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepo.GetAllComments()
	if err != nil {
		return nil, err
	}

	var response []dto.GetAllCommentResponse
	for _, comment := range comments {
		user, err := c.userRepo.GetUserByIDComment(comment.UserId)
		if err != nil {
			return nil, err
		}

		photo, err := c.photoRepo.GetPhotoByID(comment.PhotoId)
		if err != nil {
			return nil, err
		}

		commentResponse := convertToCommentResponse(comment, user, photo)
		response = append(response, commentResponse)
	}
	return response, nil
}

func (c *commentService) UpdateComment(id uint, comment *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr) {
	commentUpdate, err := c.commentRepo.GetCommentById(id)

	if err != nil {
		return nil, err
	}

	newComment := comment.CommentUpdateRequestToModel()

	updatedComment, err := c.commentRepo.UpdateComment(commentUpdate, newComment)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	response := &dto.UpdateCommentResponse{
		ID:        updatedComment.ID,
		Message:   updatedComment.Message,
		PhotoId:   updatedComment.PhotoId,
		UserId:    updatedComment.UserId,
		UpdatedAt: updatedComment.UpdatedAt,
	}

	return response, nil
}

func (c *commentService) DeleteComment(id uint) (*dto.DeleteCommentResponse, errs.MessageErr) {
	comment, err := c.commentRepo.GetCommentById(id)

	if err != nil {
		return nil, err
	}

	if err := c.commentRepo.DeleteComment(comment); err != nil {
		return nil, err
	}

	deleteResponse := &dto.DeleteCommentResponse{
		Message: "Your comment has been successfully deleted",
	}
	return deleteResponse, nil
}

func convertToCommentResponse(c models.Comment, user *models.User, photo *models.Photo) dto.GetAllCommentResponse {
	commentResponse := dto.GetAllCommentResponse{
		ID:        c.ID,
		Message:   c.Message,
		PhotoId:   c.PhotoId,
		UserId:    c.UserId,
		UpdatedAt: c.UpdatedAt,
		CreatedAt: c.CreatedAt,
		User: dto.UserDataForComment{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
		Photo: dto.PhotoDataForComment{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoURL: photo.PhotoURL,
			UserId:   photo.UserId,
		},
	}

	return commentResponse
}

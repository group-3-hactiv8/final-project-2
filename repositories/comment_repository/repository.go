package comment_repository

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
)

type CommentRepository interface {
	CreateComment(user *models.User, comment *models.Comment) (*models.Comment, errs.MessageErr)
	GetCommentById(id uint) (*models.Comment, errs.MessageErr)
	GetCommentByUserId(userId uint) ([]models.Comment, errs.MessageErr)
	GetAllComments() ([]models.Comment, errs.MessageErr)
	GetCommentByPhotoId(photoId uint) ([]models.Comment, errs.MessageErr)
	UpdateComment(Comment *models.Comment, cmtUpdate *models.Comment) (*models.Comment, errs.MessageErr)
	DeleteComment(category *models.Comment) errs.MessageErr
}
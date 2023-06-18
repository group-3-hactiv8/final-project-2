package user_repository

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
)

type UserRepository interface {
	GetUserByID(user *models.User) errs.MessageErr
	GetUserByEmail(user *models.User) errs.MessageErr
	RegisterUser(user *models.User) (*models.User, errs.MessageErr)
	LoginUser(user *models.User) errs.MessageErr
	UpdateUser(user *models.User) (*models.User, errs.MessageErr)
	DeleteUser(id uint) errs.MessageErr
	GetUserByIDComment(id uint) (*models.User, errs.MessageErr)
}

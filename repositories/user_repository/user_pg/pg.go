package user_pg

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/user_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) RegisterUser(newUser *models.User) (*models.User, errs.MessageErr) {
	if err := u.db.Create(newUser).Error; err != nil {
		log.Println(err.Error())
		message := fmt.Sprintf("Failed to register a new user with username %s", newUser.Username)
		error := errs.NewInternalServerError(message)
		return nil, error
	}

	return newUser, nil
}

func (u *userPG) LoginUser(user *models.User) errs.MessageErr {
	return nil
}

func (u *userPG) UpdateUser(user *models.User) (*models.User, errs.MessageErr) {
	return nil, nil
}

func (u *userPG) DeleteUser(id uint) errs.MessageErr {
	return nil
}

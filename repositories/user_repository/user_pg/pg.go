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

func (u *userPG) GetUserByID(user *models.User) errs.MessageErr {
	err := u.db.Where("id = ?", user.ID).Take(&user).Error
	// Karna di Take, objek user akan terupdate, termasuk passwordnya.
	// Makannya kita simpen dulu password dari request nya di service level.

	if err != nil {
		message := fmt.Sprintf("User with ID %v not found", user.ID)
		err2 := errs.NewNotFound(message)
		return err2
	}

	return nil
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
	err := u.db.Where("username = ?", user.Username).Take(&user).Error
	// Karna di Take, objek user akan terupdate, termasuk passwordnya.
	// Makannya kita simpen dulu password dari request nya di service level.

	if err != nil {
		err2 := errs.NewBadRequest("Wrong username/password")
		return err2
	}

	return nil
}

func (u *userPG) UpdateUser(user *models.User) (*models.User, errs.MessageErr) {
	initialUser := &models.User{}
	err := u.db.Where("id = ?", user.ID).Take(&initialUser).Error

	if err != nil {
		message := fmt.Sprintf("User with ID %v not found", user.ID)
		err2 := errs.NewNotFound(message)
		return nil, err2
	}

	err = u.db.Model(user).Updates(user).Error

	if err != nil {
		err2 := errs.NewBadRequest(err.Error())
		return nil, err2
	}

	user.Age = initialUser.Age

	return user, nil
}

func (u *userPG) DeleteUser(id uint) errs.MessageErr {
	initialUser := &models.User{}

	err := u.db.Where("id = ?", id).Take(&initialUser).Error

	if err != nil {
		message := fmt.Sprintf("User with ID %v not found", id)
		err2 := errs.NewNotFound(message)
		return err2
	}

	err = u.db.Model(initialUser).Delete(initialUser).Error

	if err != nil {
		err3 := errs.NewInternalServerError(err.Error())
		return err3
	}

	return nil
}

func (u *userPG) GetUserByIDComment(id uint) (*models.User, errs.MessageErr) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with id %d is not found", id))
	}

	return &user, nil
}
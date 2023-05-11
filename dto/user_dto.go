package dto

import (
	"final-project-2/models"
	"final-project-2/pkg/errs"

	"github.com/asaskevich/govalidator"
)

type NewUserRequest struct {
	Username string `json:"username" valid:"required~Your username is required"`
	Email    string `json:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string `json:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `json:"age" valid:"required~Your age is required"`
}

// stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func (u *NewUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	isAbove8 := govalidator.InRangeInt(u.Age, 9, MaxInt) // kalau 9 tetep true (lower bound)

	if !isAbove8 {
		return errs.NewUnprocessableEntity("Age must have value above 8")
	}

	return nil

}

func (u *NewUserRequest) UserRequestToModel() *models.User {
	return &models.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Age:      u.Age,
	}
}

type NewUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Age      int    `json:"age"`
}

type LoginUserRequest struct {
	Username string `json:"username" valid:"required~Your username is required"`
	Password string `json:"password" valid:"required~Your password is required"`
}

func (u *LoginUserRequest) ValidateStruct() errs.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}

func (u *LoginUserRequest) LoginUserRequestToModel() *models.User {
	return &models.User{
		Username: u.Username,
		Password: u.Password,
	}
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
}

type UpdateUserResponse struct {
}

type DeleteUserResponse struct {
}

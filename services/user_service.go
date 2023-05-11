package services

import (
	"final-project-2/dto"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/user_repository"
)

type UserService interface {
	RegisterUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	LoginUser(payload *dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr)
	UpdateUser(payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr)
	DeleteUser(id uint) (*dto.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) RegisterUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	newUser := payload.UserRequestToModel()

	createdUser, err := u.userRepo.RegisterUser(newUser)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUserResponse{
		Username: createdUser.Username,
		Email:    createdUser.Email,
		ID:       createdUser.ID,
		Age:      createdUser.Age,
	}

	return response, nil
}

func (u *userService) LoginUser(payload *dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr) {
	return nil, nil
}

func (u *userService) UpdateUser(payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr) {
	return nil, nil
}

func (u *userService) DeleteUser(id uint) (*dto.DeleteUserResponse, errs.MessageErr) {
	return nil, nil
}

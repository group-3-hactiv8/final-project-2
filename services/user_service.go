package services

import (
	"final-project-2/dto"
	"final-project-2/helpers"
	"final-project-2/pkg/errs"
	"final-project-2/repositories/user_repository"
)

type UserService interface {
	RegisterUser(payload *dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	LoginUser(payload *dto.LoginUserRequest) (*dto.LoginUserResponse, errs.MessageErr)
	UpdateUser(id int, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr)
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
	user := payload.LoginUserRequestToModel()
	passwordFromRequest := user.Password

	err := u.userRepo.LoginUser(user)

	if err != nil {
		return nil, err
	}

	isTheSame := helpers.ComparePass([]byte(user.Password), []byte(passwordFromRequest))
	// harus pake method comparePass ini instead of pake statement Where buat nyari di DB.
	// karena passwordnya disimpan setelah di hash pada function BeforeCreate.

	if !isTheSame {
		err := errs.NewBadRequest("Wrong username/password")
		return nil, err
	}

	token := helpers.GenerateToken(user.ID, user.Username)

	response := &dto.LoginUserResponse{
		Token: token,
	}

	return response, nil
}

func (u *userService) UpdateUser(id int, payload *dto.UpdateUserRequest) (*dto.UpdateUserResponse, errs.MessageErr) {
	userUpdateRequest := payload.UpdateUserRequestToModel()
	if id < 1 {
		idError := errs.NewBadRequest("ID value must be positive")
		return nil, idError
	}

	userUpdateRequest.ID = uint(id)

	updatedUser, err := u.userRepo.UpdateUser(userUpdateRequest)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateUserResponse{
		Username:  updatedUser.Username,
		Email:     updatedUser.Email,
		ID:        updatedUser.ID,
		Age:       updatedUser.Age,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(id uint) (*dto.DeleteUserResponse, errs.MessageErr) {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	response := &dto.DeleteUserResponse{
		Message: "Your account has been successfully deleted",
	}

	return response, nil
}

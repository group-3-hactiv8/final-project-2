package http_handlers

import (
	"final-project-2/dto"
	"final-project-2/pkg/errs"
	"final-project-2/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService: userService}
}

// RegisterUser godoc
//
//	@Summary		Register a user
//	@Description	Register a user by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.NewUserRequest	true	"Create user request body"
//	@Success		201		{object}	dto.NewUserResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/users/register [post]
func (u *userHandler) RegisterUser(ctx *gin.Context) {
	var requestBody dto.NewUserRequest

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	createdUser, err3 := u.userService.RegisterUser(&requestBody)
	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// LoginUser godoc
//
//	@Summary		Login
//	@Description	Login by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.LoginUserRequest	true	"Login user request body"
//	@Success		200		{object}	dto.LoginUserResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/users/login [post]
func (u *userHandler) LoginUser(ctx *gin.Context) {
	var requestBody dto.LoginUserRequest

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	tokenResponse, err3 := u.userService.LoginUser(&requestBody)

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, tokenResponse)
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update a user by json
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.UpdateUserRequest	true	"Update a user request body"
//	@Success		200		{object}	dto.UpdateUserResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/users [put]
func (u *userHandler) UpdateUser(ctx *gin.Context) {
	// mustget = ambil data dari middleware authentication.
	// Tp hasil returnnya hanya empty interface, jadi harus
	// di cast dulu ke jwt.MapClaims.
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	var requestBody dto.UpdateUserRequest

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := requestBody.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	updatedUserResponse, err3 := u.userService.UpdateUser(userId, &requestBody)

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, updatedUserResponse)

}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by JWT from header
//	@Tags			users
//	@Produce		json
//	@Success		200		{object}	dto.DeleteUserResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Router			/users [delete]
func (u *userHandler) DeleteUser(ctx *gin.Context) {
	// mustget = ambil data dari middleware authentication.
	// Tp hasil returnnya hanya empty interface, jadi harus
	// di cast dulu ke jwt.MapClaims.
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	response, err2 := u.userService.DeleteUser(userId)

	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

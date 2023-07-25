package http_handlers

import (
	"final-project-2/dto"
	"final-project-2/pkg/errs"
	"final-project-2/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService services.SocialMediaService) *socialMediaHandler {
	return &socialMediaHandler{socialMediaService: socialMediaService}
}

// CreateSocialMedia godoc
//
//	@Summary		Create a social media
//	@Description	Create a social media by json
//	@Tags			socialmedias
//	@Accept			json
//	@Produce		json
//	@Param			socialmedia	body		dto.NewSocialMediaRequest	true	"Create a social media request body"
//	@Success		201		{object}	dto.NewSocialMediaResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/socialmedias [post]
func (sm *socialMediaHandler) CreateSocialMedia(ctx *gin.Context) {
	// mustget = ambil data dari middleware authentication.
	// Tp hasil returnnya hanya empty interface, jadi harus
	// di cast dulu ke jwt.MapClaims.
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	var requestBody dto.NewSocialMediaRequest

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

	newSMResponse, err3 := sm.socialMediaService.CreateSocialMedia(&requestBody, userId)

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, newSMResponse)
}

// ViewAllSocialMedia godoc
//
//	@Summary		View all social medias
//	@Description	View all social medias by json
//	@Tags			socialmedias
//	@Produce		json
//	@Success		200		{object}	dto.AllSocialMediasResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/socialmedias [get]
func (sm *socialMediaHandler) GetAllSocialMedias(ctx *gin.Context) {
	allSMResponse, err3 := sm.socialMediaService.GetAllSocialMedias()

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, allSMResponse)
}

// UpdateSocialMedia godoc
//
//	@Summary		Update a Social Media
//	@Description	Update a Social Media by json
//	@Tags			socialmedias
//	@Accept			json
//	@Produce		json
//	@Param			socialMedia	body		dto.NewSocialMediaRequest	true	"Update a social media request body"
//	@Param			socialMediaId		path		uint					true	"Social Media ID request"
//	@Success		200		{object}	dto.UpdateSocialMediaResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/socialmedias/{socialMediaId} [put]
func (sm *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	// ambil socialmedia id dari path variable
	id, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	var requestBody dto.NewSocialMediaRequest

	err = ctx.ShouldBindJSON(&requestBody)
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

	updatedSMResponse, err3 := sm.socialMediaService.UpdateSocialMedia(id, &requestBody)

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, updatedSMResponse)
}

// DeleteSocialMedia godoc
//
//	@Summary		Delete a Social Media
//	@Description	Delete a Social Media by param
//	@Tags			socialmedias
//	@Produce		json
//	@Param			socialMediaId		path		uint					true	"Social Media ID request"
//	@Success		200		{object}	dto.DeleteSocialMediaResponse
// @Param Authorization header string true "Insert your access token" default(Bearer <Add your access token here>)
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/socialmedias/{socialMediaId} [delete]
func (sm *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
	// ambil socialmedia id dari path variable
	id, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	deletedSMResponse, err3 := sm.socialMediaService.DeleteSocialMedia(id)

	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, deletedSMResponse)
}

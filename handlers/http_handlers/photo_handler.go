package http_handlers

import (
	"final-project-2/dto"
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"final-project-2/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService services.PhotoService
}

func NewPhotoHandler(photoService services.PhotoService) *photoHandler {
	return &photoHandler{photoService: photoService}
}

func (p *photoHandler) CreatePhoto(ctx *gin.Context) {
	// userData, ok := ctx.MustGet("userData").(*models.User)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	username := userData["username"].(string)
	user := &models.User{
		Username: username,
	}
	user.ID = userId
	// if !ok {
	// 	fmt.Println(ctx.MustGet("userData"))
	// 	newError := errs.NewBadRequest("Failed to get user data")
	// 	ctx.JSON(newError.StatusCode(), newError)
	// 	return
	// }
	var requestBody dto.CreatePhotoRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdPhoto, err := p.photoService.CreatePhoto(user, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdPhoto)
}

func (p *photoHandler) GetAllPhotos(ctx *gin.Context) {
	_, ok := ctx.MustGet("userData").(*models.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	photos, err := p.photoService.GetAllPhotos()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (p *photoHandler) UpdatePhoto(ctx *gin.Context) {
	var requestBody dto.UpdatePhotoRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	photoID := ctx.Param("photoID")
	photoIDUint, err := strconv.ParseUint(photoID, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Photo id should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	updatedPhoto, err2 := p.photoService.UpdatePhoto(uint(photoIDUint), &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}

func (p *photoHandler) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoID")
	photoIDUint, err := strconv.ParseUint(photoID, 10, 32)
	if err != nil {
		newError := errs.NewBadRequest("Photo id should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err2 := p.photoService.DeletePhoto(uint(photoIDUint))
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

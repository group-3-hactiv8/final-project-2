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

// CreatePhoto godoc
//
//	@Summary		Create a photo
//	@Description	Create a photo by json
//	@Tags			photos
//	@Accept			json
//	@Produce		json
//	@Param			comment	body		dto.CreatePhotoRequest	true	"Create a photo request body"
//	@Success		201		{object}	dto.CreatePhotoResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/photos [post]
func (p *photoHandler) CreatePhoto(ctx *gin.Context) {
	// userData, ok := ctx.MustGet("userData").(*models.User)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	username := userData["username"].(string)
	user := &models.User{
		Username: username,
	}
	user.ID = userId
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

// GetAllComment godoc
//
//	@Summary		Get all photos
//	@Description	Get all photos by json
//	@Tags			photos
//	@Produce		json
//	@Success		200		{object}	dto.GetAllPhotosResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/ [get]
func (p *photoHandler) GetAllPhotos(ctx *gin.Context) {

	photos, err := p.photoService.GetAllPhotos()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

// UpdatePhoto godoc
//
//	@Summary		Update a Photo
//	@Description	Update a Photo by json
//	@Tags			photos
//	@Accept			json
//	@Produce		json
//	@Param			photo	body		dto.UpdatePhotoRequest	true	"Update a photos request body"
//	@Param			commentId		path		uint					true	"photos ID request"
//	@Success		200		{object}	dto.UpdatePhotoResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/photos/{photosId} [put]
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

// DeletePhoto godoc
// @Summary Delete a photo
// @Description Delete a specific photo by ID
// @Tags photos
// @Param photoID path int true "Photo ID"
// @Produce json
// @Success 200 {object} dto.DeletePhotoResponse
// @Failure 400 {object} errs.MessageErrData
// @Failure 404 {object} errs.MessageErrData
// @Failure 500 {object} errs.MessageErrData
// @Router /photos/{photoID} [delete]
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

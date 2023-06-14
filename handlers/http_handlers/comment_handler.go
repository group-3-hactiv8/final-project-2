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

type commentHandler struct {
	commentService services.CommentService
}

func NewCommentHandler(commentService services.CommentService) *commentHandler {
	return &commentHandler{commentService: commentService}
}

// CreateComment godoc
//
//	@Summary		Create a comment
//	@Description	Create a comment by json
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			comment	body		dto.NewCommentRequest	true	"Create a comment request body"
//	@Success		201		{object}	dto.NewCommentResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/comment [post]
func (c *commentHandler) CreateComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	username := userData["username"].(string)
	user := &models.User{
		Username: username,
	}
	user.ID = userId

	var requestBody dto.NewCommentRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdComment, err := c.commentService.CreateComment(user, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdComment)
}

// GetCommentByUserId godoc
//
//	@Summary		View all comments of a user
//	@Description	View all comments of a user by json
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			userId		path		uint					true	"user ID request"
//	@Success		200		{object}	dto.GetAllCommentResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/comment/user/{userId} [get]
func (c *commentHandler) GetCommentByUserId(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	username := userData["username"].(string)
	user := &models.User{
		Username: username,
	}
	user.ID = userId

	comments, err := c.commentService.GetCommentByUserId(userId)

	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// GetCommentByPhotoId godoc
//
//	@Summary		View all comments of a photo
//	@Description	View all comments of a photo by json
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			photoId		path		uint					true	"photo  ID request"
//	@Success		200		{object}	dto.GetAllCommentResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/comment/photo/{photoId} [get]
func (c *commentHandler) GetCommentByPhotoId(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid photoId",
		})
		return
	}

	comments, err := c.commentService.GetCommentByPhotoId(uint(photoId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	// Mengembalikan respons sukses dengan data komentar
	ctx.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

// ViewAllComment godoc
//
//	@Summary		View all comment
//	@Description	View all comment by json
//	@Tags			comment
//	@Produce		json
//	@Success		200		{object}	dto.GetAllCommentResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/comment [get]
func (c *commentHandler) GetAllComment(ctx *gin.Context) {
	comments, err := c.commentService.GetAllComment()
	if err != nil {
		// Mengembalikan respons error jika terjadi kesalahan
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	// Mengembalikan respons sukses dengan data komentar
	ctx.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

// UpdateComment godoc
//
//	@Summary		Update a Comment
//	@Description	Update a Comment by json
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Param			comment	body		dto.UpdateCommentRequest	true	"Update a comment request body"
//	@Param			commentId		path		uint					true	"comment ID request"
//	@Success		200		{object}	dto.UpdateCommentResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/comment/{commentId} [put]
func (c *commentHandler) UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")
	commentIDUint, err := strconv.ParseUint(commentID, 10, 16)
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	var updateRequest dto.UpdateCommentRequest
	err = ctx.ShouldBindJSON(&updateRequest)
	if err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	err2 := updateRequest.ValidateStruct()
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	updateComment, err3 := c.commentService.UpdateComment(uint(commentIDUint), &updateRequest)
	if err3 != nil {
		ctx.JSON(err3.StatusCode(), err3)
		return
	}

	ctx.JSON(http.StatusOK, updateComment)
}

// DeleteComment godoc
//
//	@Summary		Delete a comment
//	@Description	Delete a comment by param
//	@Tags			comment
//	@Produce		json
//	@Param			commentId		path		uint					true	"comment ID request"
//	@Success		200		{object}	dto.DeleteCommentResponse
//	@Failure		401		{object}	errs.MessageErrData
//	@Failure		400		{object}	errs.MessageErrData
//	@Router			/comment/{commentId} [delete]
func (c *commentHandler) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	commentIdUint, err := strconv.ParseUint(commentId, 10, 16)
	if err != nil {
		idError := errs.NewBadRequest("Invalid ID format")
		ctx.JSON(idError.StatusCode(), idError)
		return
	}

	response, err2 := c.commentService.DeleteComment(uint(commentIdUint))

	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

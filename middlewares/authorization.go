package middlewares

import (
	"errors"
	"final-project-2/database"
	"final-project-2/models"
	"final-project-2/pkg/errs"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// func ProductAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetDB()
// 		// misal ada user yg mau akses product berdasarkan productId
// 		// nya melalui param. Param = path variable
// 		productId, err := strconv.Atoi(c.Param("productId"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error":   "Bad Request",
// 				"message": "Invalid parameter (gada productId)",
// 			})
// 			return
// 		}
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		userId := uint(userData["id"].(float64))
// 		Product := models.Product{}

// 		// cek product yg dicari berdasarkan product id nya ada atau engga
// 		err = db.Select("user_id").First(&Product, uint(productId)).Error

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Not Found",
// 				"message": "Product not found",
// 			})
// 			return
// 		}

// 		// product nya ada, tp cek dulu userId nya sama dengan
// 		// userId si user yg lg login ngga?
// 		if Product.UserId != userId {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this product",
// 			})
// 			return
// 		}

// 		c.Next()
// 	}
// }

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()
		// misal ada user yg mau akses product berdasarkan socialMediaId
		// nya melalui param. Param = path variable
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
		if err != nil {
			badRequestError := errs.NewBadRequest("Invalid parameter for socialMediaId")
			c.AbortWithStatusJSON(badRequestError.StatusCode(), badRequestError)
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		requestedSocialMedia := &models.SocialMedia{}

		// cek socialmedia yg dicari berdasarkan socialmedia id nya ada atau engga
		err = db.Where("id = ?", socialMediaId).Take(&requestedSocialMedia).Error

		if err != nil {
			notFoundError := errs.NewNotFound("Social Media not found")
			c.AbortWithStatusJSON(notFoundError.StatusCode(), notFoundError)
			return
		}

		// socialmedia nya ada, tp cek dulu userId nya sama dengan
		// userId si user yg lg login ngga?
		if requestedSocialMedia.UserId != userId {
			unauthorizedError := errs.NewUnauthorized("You are not allowed to access this Social Media")
			c.AbortWithStatusJSON(unauthorizedError.StatusCode(), unauthorizedError)
			return
		}

		c.Next()
	}

}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()

		photoID, err := strconv.Atoi(c.Param("photoID"))
		if err != nil {
			badRequestError := fmt.Errorf("Invalid photo ID: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": badRequestError.Error()})
			return
		}

		// userData, ok := c.Get("userData").(jwt.MapClaims)
		userData, ok := c.MustGet("userData").(jwt.MapClaims)
		if !ok {
			unauthorizedError := errors.New("Unauthorized access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorizedError.Error()})
			return
		}
		userID := uint(userData["id"].(float64))
		requestedPhoto := &models.Photo{}

		err = db.First(requestedPhoto, "id = ?", photoID).Error
		if err != nil {
			notFoundError := fmt.Errorf("Photo not found: %v", err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": notFoundError.Error()})
			return
		}

		if requestedPhoto.UserId != userID {
			unauthorizedError := errors.New("You are not allowed to access this Photo")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": unauthorizedError.Error()})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetPostgresInstance()
		// misal ada user yg mau akses product berdasarkan commentId
		// nya melalui param. Param = path variable
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			badRequestError := errs.NewBadRequest("Invalid parameter for commentId")
			c.AbortWithStatusJSON(badRequestError.StatusCode(), badRequestError)
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		requestedComment := &models.Comment{}
		
		err = db.Where("id = ?", commentId).Take(&requestedComment).Error

		if err != nil {
			notFoundError := errs.NewNotFound("Comment not found")
			c.AbortWithStatusJSON(notFoundError.StatusCode(), notFoundError)
			return
		}

		if requestedComment.UserId != userId {
			unauthorizedError := errs.NewUnauthorized("You are not allowed to access this Comment")
			c.AbortWithStatusJSON(unauthorizedError.StatusCode(), unauthorizedError)
			return
		}

		c.Next()
	}

}
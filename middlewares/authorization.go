package middlewares

import (
	"final-project-2/database"
	"final-project-2/models"
	"final-project-2/pkg/errs"
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

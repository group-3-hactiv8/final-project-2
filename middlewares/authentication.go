package middlewares

import (
	"final-project-2/database"
	"final-project-2/helpers"
	"final-project-2/models"
	"final-project-2/repositories/user_repository/user_pg"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken // dapet MapClaims yg isinya id dan email

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		data := verifyToken.(jwt.MapClaims)

		id := uint(data["id"].(float64))
		// ada key "username" ga?
		if _, isExist := data["username"]; !isExist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "invalid token",
			})
			return
		}
		username := data["username"].(string)

		initialUser := &models.User{}
		initialUser.ID = id

		db := database.GetPostgresInstance()
		userRepo := user_pg.NewUserPG(db)

		// does a user exist with this id?
		if errNotFound := userRepo.GetUserByID(initialUser); errNotFound != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "invalid token",
			})
			return
		}

		// if exist, is the username is the same with the one from the token?
		if initialUser.Username != username {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": "invalid token",
			})
			return
		}

		// nyimpen claim untuk diambil di endpoint berikutnya
		c.Set("userData", verifyToken)
		c.Next() // lanjut ke endpoint berikutnya, yakni ke controller
	}
}

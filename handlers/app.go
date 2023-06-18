package handlers

import (
	"final-project-2/database"
	_ "final-project-2/docs"
	"final-project-2/handlers/http_handlers"
	"final-project-2/middlewares"
	"final-project-2/repositories/comment_repository/comment_pg"
	"final-project-2/repositories/photo_repository/photo_pg"
	"final-project-2/repositories/social_media_repository/social_media_pg"
	"final-project-2/repositories/user_repository/user_pg"
	"final-project-2/services"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
)

// @title MyGram API
// @version 1.0
// @description This is a server for MyGram Application.
// @termsOfService http://swagger.io/terms/
// @contact.name Swagger API Team
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	database.StartDB()
	db := database.GetPostgresInstance()

	router := gin.Default()

	router.GET("/health-check", func (c *gin.Context){
		c.JSON(200, gin.H{
			"appName" : "MyGramApp",
		})
	})

	userRepo := user_pg.NewUserPG(db)
	userService := services.NewUserService(userRepo)
	userHandler := http_handlers.NewUserHandler(userService)

	usersRouter := router.Group("/users")
	{
		usersRouter.POST("/register", userHandler.RegisterUser)
		usersRouter.POST("/login", userHandler.LoginUser)
		usersRouter.PUT("/", middlewares.Authentication(), userHandler.UpdateUser)
		usersRouter.DELETE("/", middlewares.Authentication(), userHandler.DeleteUser)
	}

	socialMediaRepo := social_media_pg.NewSocialMediaPG(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaHandler := http_handlers.NewSocialMediaHandler(socialMediaService)

	socialMediasRouter := router.Group("/socialmedias")
	socialMediasRouter.Use(middlewares.Authentication())
	{
		socialMediasRouter.POST("/", socialMediaHandler.CreateSocialMedia)
		socialMediasRouter.GET("/", socialMediaHandler.GetAllSocialMedias)
		socialMediasRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), socialMediaHandler.UpdateSocialMedia)
		socialMediasRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), socialMediaHandler.DeleteSocialMedia)
	}

	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := services.NewPhotoService(photoRepo, userRepo)
	photoHandler := http_handlers.NewPhotoHandler(photoService)

	photoRouter := router.Group("/photos")
	photoRouter.Use(middlewares.Authentication())
	{
		photoRouter.POST("/", photoHandler.CreatePhoto)
		photoRouter.GET("/", photoHandler.GetAllPhotos)
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), photoHandler.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), photoHandler.DeletePhoto)
	}

	commentRepo := comment_pg.NewCommentPG(db)
	commentService := services.NewCommentService(commentRepo, photoRepo, userRepo)
	commentHandler := http_handlers.NewCommentHandler(commentService)

	commentRouter := router.Group("/comment")
	commentRouter.Use(middlewares.Authentication())
	{
		commentRouter.POST("/", commentHandler.CreateComment)
		commentRouter.GET("/", commentHandler.GetAllComment)
		commentRouter.GET("/user/:userId", commentHandler.GetCommentByUserId)
		commentRouter.GET("/photo/:photoId", commentHandler.GetCommentByPhotoId)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), commentHandler.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), commentHandler.DeleteComment)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router

}

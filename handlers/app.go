package handlers

import (
	"final-project-2/database"
	// "final-project-2/docs"
	"final-project-2/handlers/http_handlers"
	"final-project-2/middlewares"
	"final-project-2/repositories/photo_repository/photo_pg"
	"final-project-2/repositories/social_media_repository/social_media_pg"
	"final-project-2/repositories/user_repository/user_pg"
	"final-project-2/services"

	"github.com/gin-gonic/gin"
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

	userRepo := user_pg.NewUserPG(db)
	userService := services.NewUserService(userRepo)
	userHandler := http_handlers.NewUserHandler(userService)

	usersRouter := router.Group("/users")
	{
		usersRouter.POST("/register", userHandler.RegisterUser)
		usersRouter.POST("/login", userHandler.LoginUser)
		usersRouter.PUT("/:id", middlewares.Authentication(), userHandler.UpdateUser)
		usersRouter.DELETE("/", middlewares.Authentication(), userHandler.DeleteUser)
	}

	// userRepo := user_pg.NewUserPG(db)
	// userService := service.NewUserService(userRepo)
	// userHandler := http_handler.NewUserHandler(userService)

	// photosRouter := router.Group("/photos")
	// {
	// 	photosRouter.POST("/", photoHandler.CreatePhoto)
	// 	photosRouter.GET("/", photoHandler.GetAllPhotos)
	// 	photosRouter.PUT("/:id", photoHandler.UpdatePhoto)
	// 	photosRouter.DELETE("/:id", photoHandler.DeletePhoto)
	// }

	// userRepo := user_pg.NewUserPG(db)
	// userService := service.NewUserService(userRepo)
	// userHandler := http_handler.NewUserHandler(userService)

	// commentsRouter := router.Group("/comments")
	// {
	// 	commentsRouter.POST("/", commentHandler.CreateComment)
	// 	commentsRouter.GET("/", commentHandler.GetAllComments)
	// 	commentsRouter.PUT("/:id", commentHandler.UpdateComment)
	// 	commentsRouter.DELETE("/:id", commentHandler.DeleteComment)
	// }

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

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router

}

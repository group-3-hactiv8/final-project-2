package handlers

import (
	"final-project-2/database"
	// _ "final-project-2/docs"
	"final-project-2/handlers/http_handlers"
	"final-project-2/repositories/user_repository/user_pg"
	"final-project-2/services"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title MyGram API
// @version 1.0
// @description This is a sample server for ... .
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
		usersRouter.PUT("/:id", userHandler.UpdateUser)
		usersRouter.DELETE("/:id", userHandler.DeleteUser)
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

	// userRepo := user_pg.NewUserPG(db)
	// userService := service.NewUserService(userRepo)
	// userHandler := http_handler.NewUserHandler(userService)

	// socialMediasRouter := router.Group("/socialMedias")
	// {
	// 	socialMediasRouter.POST("/", socialMediaHandler.CreateSocialMedia)
	// 	socialMediasRouter.GET("/", socialMediaHandler.GetAllSocialMedias)
	// 	socialMediasRouter.PUT("/:id", socialMediaHandler.UpdateSocialMedia)
	// 	socialMediasRouter.DELETE("/:id", socialMediaHandler.DeleteSocialMedia)
	// }

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router

}

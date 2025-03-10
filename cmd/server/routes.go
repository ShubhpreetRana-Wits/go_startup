package server

import (
	"example.com/startup/internal/dtos"
	"example.com/startup/internal/infrastucture/repository"
	"example.com/startup/internal/interfaces/controllers"
	"example.com/startup/internal/interfaces/middleware"
	"example.com/startup/internal/interfaces/usecases"
	"example.com/startup/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app fiber.Router, db *db.PostgresClient) {

	// Create repositories
	genrateUrlRepo := repository.NewGenerateUrlRepository(db)

	// Create use cases
	generateUrlUseCase := usecases.NewGenrateUrlUsecase(genrateUrlRepo)

	// Create controllers
	genrateUrlController := controllers.NewGenerateURLController(generateUrlUseCase)

	appGroup := app.Group("/api/v1/generateUrl")
	appGroup.Post("/", middleware.ValidateParams[dtos.GenerateUrlRequest](), genrateUrlController.GenerateURL)
	app.Get("/", middleware.JWTUrlMiddleware(), middleware.JWTMetadataMiddleware(), genrateUrlController.GetURL)
	app.Delete("/", middleware.JWTUrlMiddleware(), middleware.JWTMetadataMiddleware(), genrateUrlController.DeleteURL)

}

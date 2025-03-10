package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"example.com/startup/pkg/config"
	"example.com/startup/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func StartServer() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbClient, err := db.NewGormDBConfig(cfg)
	if err != nil {
		fmt.Println("Error in connecting to database")
	}
	defer dbClient.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	InitializeRoutes(app, dbClient)
	// Graceful shutdown
	go func() {
		if err := app.Listen(":" + cfg.ServerPort); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server gracefully stopped")
}

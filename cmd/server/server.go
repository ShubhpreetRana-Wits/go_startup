package server

import (
	"fmt"
	"log"
	"net"

	"example.com/startup/grpc/controller"
	pb "example.com/startup/grpc/generated"
	"example.com/startup/internal/infrastucture/repository"
	"example.com/startup/internal/interfaces/usecases"
	"example.com/startup/pkg/config"
	"example.com/startup/pkg/db"
	"google.golang.org/grpc"
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

	// app := fiber.New(fiber.Config{
	// 	ErrorHandler: func(c *fiber.Ctx, err error) error {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error": err.Error(),
	// 		})
	// 	},
	// })

	// InitializeRoutes(app, dbClient)

	// // Graceful shutdown
	// go func() {
	// 	if err := app.Listen(":" + cfg.ServerPort); err != nil {
	// 		log.Fatalf("Failed to start server: %v", err)
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutting down server...")

	// if err := app.Shutdown(); err != nil {
	// 	log.Fatalf("Server shutdown failed: %v", err)
	// }
	// log.Println("Server gracefully stopped")

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// Create repositories
	genrateUrlRepo := repository.NewGenerateUrlRepository(dbClient)

	// Create use cases
	generateUrlUseCase := usecases.NewGenrateUrlUsecase(genrateUrlRepo)

	serverController := controller.NewServer(generateUrlUseCase)
	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterGenerateURLServiceServer(grpcServer, serverController)

	fmt.Println("gRPC Server is running on port 50051...")

	// Start serving requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

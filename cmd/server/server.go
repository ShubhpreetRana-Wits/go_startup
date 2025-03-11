package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"example.com/startup/grpc/client"
	"example.com/startup/grpc/controller"
	pb "example.com/startup/grpc/generated"
	"example.com/startup/internal/infrastucture/db"
	"example.com/startup/internal/infrastucture/repository"
	"example.com/startup/internal/interfaces/usecases"
	"example.com/startup/pkg/config"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func StartServer() {
	cfg := loadConfiguration()
	dbClient := setupDatabase(cfg)
	defer dbClient.Close()

	app := setupHTTPServer(cfg, dbClient)
	go startHTTPServer(app, cfg)

	grpcServer, listener := startGRPCServer(dbClient, cfg)

	handleGracefulShutdown(app, grpcServer, listener)
}

func loadConfiguration() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}

func setupDatabase(cfg *config.Config) *db.PostgresClient {
	dbClient, err := db.NewGormDBConfig(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return dbClient
}

func setupHTTPServer(cfg *config.Config, dbClient *db.PostgresClient) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	InitializeRoutes(app, dbClient)
	return app
}

func startHTTPServer(app *fiber.App, cfg *config.Config) {
	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func startGRPCServer(dbClient *db.PostgresClient, cfg *config.Config) (*grpc.Server, net.Listener) {
	listener, err := net.Listen("tcp", ":"+cfg.GRPC_PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	generateUrlRepo := repository.NewGenerateUrlRepository(dbClient)
	generateUrlUseCase := usecases.NewGenrateUrlUsecase(generateUrlRepo)
	serverController := controller.NewServer(generateUrlUseCase, client.SendMessageClient{Cfg: cfg})

	grpcServer := grpc.NewServer()
	pb.RegisterGenerateURLServiceServer(grpcServer, serverController)

	go func() {
		log.Println("gRPC Server is running on port 50051...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	return grpcServer, listener
}

func handleGracefulShutdown(app *fiber.App, grpcServer *grpc.Server, listener net.Listener) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}
	log.Println("HTTP server gracefully stopped")

	grpcServer.GracefulStop()
	if err := listener.Close(); err != nil {
		log.Fatalf("Failed to close gRPC listener: %v", err)
	}
	log.Println("gRPC server gracefully stopped")
}

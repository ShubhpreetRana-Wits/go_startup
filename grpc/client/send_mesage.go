package client

import (
	"context"
	"log"
	"time"

	services "example.com/startup/grpc/generated"
	"example.com/startup/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SendMessageClient struct {
	Cfg *config.Config
}

func (s *SendMessageClient) SendMessage() (*services.SentMessageResponse, error) {
	log.Println("Initializing gRPC connection...")

	// Set up connection to gRPC server
	conn, err := grpc.NewClient(s.Cfg.NOTI_GRPC_HOST+":"+s.Cfg.NOTI_GRPC_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server at %s:%s - Error: %v", s.Cfg.NOTI_GRPC_HOST, s.Cfg.NOTI_GRPC_PORT, err)
	}
	defer func() {
		log.Println("Closing gRPC connection...")
		conn.Close()
	}()

	client := services.NewNotificationServiceClient(conn)
	log.Println("gRPC client initialized successfully.")

	// Create request
	req := &services.ConsumeMessageRequest{
		User:    "Shobit",
		Message: "Url Generated Success",
		Medium:  "mail",
	}
	log.Printf("Sending message request: %+v\n", req)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call SendMessage function
	resp, err := client.SendMessage(ctx, req)
	if err != nil {
		log.Printf("Error calling SendMessage: %v\n", err)
		return nil, err
	}

	log.Printf("Received response: %+v\n", resp)
	return resp, nil
}

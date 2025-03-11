package controller

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"example.com/startup/grpc/client"
	pb "example.com/startup/grpc/generated"
	"example.com/startup/internal/domain/entities"
	"example.com/startup/internal/dtos"
	"example.com/startup/internal/interfaces/usecases"
	"example.com/startup/pkg/config"
	"example.com/startup/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server struct implements the gRPC service
type Server struct {
	pb.UnimplementedGenerateURLServiceServer
	generateUrlUseCase usecases.GeneratedUrlUseCase
	messageClient      client.SendMessageClient
}

// NewServer initializes a new gRPC Server instance
func NewServer(generateUrlUseCase usecases.GeneratedUrlUseCase, messageClient client.SendMessageClient) *Server {
	return &Server{
		generateUrlUseCase: generateUrlUseCase,
		messageClient:      messageClient,
	}
}

// GenerateURL handles URL generation request
func (s *Server) GenerateURL(ctx context.Context, req *pb.GenerateUrlRequest) (*pb.GenerateUrlResponse, error) {
	log.Println("Received GenerateURL request:", req)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return nil, status.Errorf(codes.Internal, "configuration error")
	}

	requestData := entities.GeneratedUrl{
		UserID:      req.UserId,
		RequestType: req.RequestType,
	}

	response, errDb := s.generateUrlUseCase.SaveUrlRequest(&requestData)
	if errDb != nil {
		log.Printf("Error saving URL data: %v", errDb)
		return nil, status.Errorf(codes.Internal, "failed to save URL request")
	}

	token, err := s.GenerateTokenWithClaim(response.ID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		_ = s.generateUrlUseCase.DeleteUrlRequest(response.ID)
		return nil, status.Errorf(codes.Internal, "failed to generate token")
	}

	redirectURL := fmt.Sprintf("%s/%s?token=%s", cfg.REDIRECTION_URL, response.ID, token)

	s.messageClient.SendMessage()
	return &pb.GenerateUrlResponse{
		RedirectUrl: redirectURL,
	}, nil
}

// GetURL handles URL retrieval request
func (s *Server) GetURL(ctx context.Context, req *pb.GetUrlRequest) (*pb.GenerateUrlInfoResponse, error) {
	log.Println("Received GetURL request:", req)

	metaData, err := s.getMetadata(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid or expired token")
	}

	response, errDb := s.generateUrlUseCase.GetUrlRequest(*metaData)
	if errDb != nil {
		log.Printf("Error retrieving URL: %v", errDb)
		return nil, status.Errorf(codes.NotFound, "URL not found")
	}

	return &pb.GenerateUrlInfoResponse{
		RequestType: response.RequestType,
		UserId:      response.UserID,
	}, nil
}

// DeleteURL handles URL deletion request
func (s *Server) DeleteURL(ctx context.Context, req *pb.DeleteUrlRequest) (*pb.BaseResponse, error) {
	log.Println("Received DeleteURL request:", req)

	metaData, err := s.getMetadata(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid or expired token")
	}

	errDb := s.generateUrlUseCase.DeleteUrlRequest(*metaData)
	if errDb != nil {
		log.Printf("Error deleting URL: %v", errDb)
		return nil, status.Errorf(codes.Internal, "failed to delete URL")
	}

	return &pb.BaseResponse{
		Message: "URL Deleted Successfully",
		Success: true,
		Code:    200,
	}, nil
}

// GenerateTokenWithClaim encrypts request data and generates a JWT
func (s *Server) GenerateTokenWithClaim(id string) (string, error) {
	log.Println("Generating token with claim...")

	encryptedData, err := utils.Encrypt(id, utils.JWTSecret)
	if err != nil {
		log.Printf("Error encrypting data: %v", err)
		return "", status.Errorf(codes.Internal, "failed to encrypt data")
	}

	claims := s.createClaims(encryptedData)
	token, err := utils.GenerateToken(claims)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", status.Errorf(codes.Internal, "failed to generate token")
	}

	return token, nil
}

// createClaims constructs JWT claims.
func (s *Server) createClaims(data string) dtos.Claims {
	return dtos.Claims{
		Data:             data,
		RegisteredClaims: utils.GenerateRegisteredClaims("your-app-name", time.Hour),
	}
}

// getMetadata extracts metadata from a JWT token
func (s *Server) getMetadata(tokenString string) (*string, error) {
	if tokenString == "" {
		log.Println("Missing token in query parameters")
		return nil, status.Errorf(codes.Unauthenticated, "token is required")
	}

	if valid := utils.ValidateToken(tokenString); !valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid or expired token")
	}

	claims, err := utils.ExtractClaims(tokenString)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to extract claims")
	}

	metadata, ok := claims["data"].(string)
	if !ok {
		return nil, status.Errorf(codes.Internal, "invalid claims format")
	}

	decryptedData, err := utils.Decrypt(metadata, utils.JWTSecret)
	if err != nil {
		log.Printf("Error decrypting metadata: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to decrypt metadata")
	}

	decryptedData = strings.ReplaceAll(decryptedData, "\x00", "")

	return &decryptedData, nil
}

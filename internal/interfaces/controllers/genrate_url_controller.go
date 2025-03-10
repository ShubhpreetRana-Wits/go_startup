package controllers

import (
	"log"
	"time"

	"example.com/startup/internal/domain/entities"
	"example.com/startup/internal/dtos"
	"example.com/startup/internal/interfaces/usecases"
	"example.com/startup/pkg/config"
	"example.com/startup/utils"
	"github.com/gofiber/fiber/v2"
)

func NewGenerateURLController(generateUrlUseCase usecases.GeneratedUrlUseCase) *GenerateURLController {
	log.Println("Initializing GenerateURLHandler...")
	return &GenerateURLController{
		generateUrlUseCase: generateUrlUseCase}
}

type GenerateURLController struct {
	generateUrlUseCase usecases.GeneratedUrlUseCase
}

func (g *GenerateURLController) GenerateURL(c *fiber.Ctx) error {
	log.Println("Received request to generate URL...")

	requestBody := c.Locals("contextData").(*dtos.GenerateUrlRequest)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	requestData := entities.GeneratedUrl{
		UserID:      requestBody.UserID,
		RequestType: requestBody.RequestType,
	}

	// Saving URL data to storage
	response, errDb := g.generateUrlUseCase.SaveUrlRequest(&requestData)
	if errDb != nil {
		log.Printf("Error saving URL data to storage: %v", errDb)
		return dtos.NewResponse(c, errDb.Error(), fiber.StatusBadRequest, "")
	}

	// Generating token
	token, err := g.GenerateTokenWithClaim(response.ID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		err := g.generateUrlUseCase.DeleteUrlRequest(response.ID)
		if err != nil {
			return dtos.NewResponse(c, err.Error(), fiber.StatusBadRequest, err)
		}
		return dtos.NewResponse(c, "Error generating token", fiber.StatusBadRequest, err)
	}
	log.Printf("Generated token: %s", token)

	return dtos.NewResponse(c, "URL generated successfully", fiber.StatusOK, response.ToResponseDTO(token, cfg.REDIRECTION_URL))
}

func (g *GenerateURLController) GetURL(c *fiber.Ctx) error {
	log.Println("Received request to retrieve URL...")

	metadata := c.Locals("jwtdata").(string)

	// Retrieve URL data from the repository using the extracted metadata
	response, errDb := g.generateUrlUseCase.GetUrlRequest(metadata)
	if errDb != nil {
		log.Printf("Error retrieving URL data from repository: %v", errDb)
		return dtos.NewResponse(c, errDb.Error(), int(fiber.StatusBadRequest), "")
	}

	log.Printf("Successfully retrieved URL data: %+v", response)

	// Return success response
	log.Println("Successfully retrieved URL")
	return dtos.NewResponse(c, "URL retrieved successfully", fiber.StatusOK, response.ToResponseInfoDTO())
}

func (g *GenerateURLController) DeleteURL(c *fiber.Ctx) error {
	log.Println("Received request to delete URL...")

	metadata := c.Locals("jwtdata").(string)
	// Delete URL data from the repository using the extracted metadata
	errDb := g.generateUrlUseCase.DeleteUrlRequest(metadata)
	if errDb != nil {
		log.Printf("Error deleting URL data from repository: %v", errDb)
		return dtos.NewResponse(c, errDb.Error(), int(fiber.StatusBadRequest), "")
	}

	// Return success response
	log.Println("Successfully deleted URL")
	return dtos.NewResponse(c, "URL deleted successfully", fiber.StatusOK, "")
}

// GenerateTokenWithClaim encrypts request data, generates a JWT, and saves the request.
func (g *GenerateURLController) GenerateTokenWithClaim(id string) (string, error) {
	log.Println("Generating token with claim...")

	encryptedData, err := utils.Encrypt(id, utils.JWTSecret)
	if err != nil {
		log.Printf("Error encrypting data: %v", err)
		return "", err
	}

	claims := g.createClaims(encryptedData)
	token, err := utils.GenerateToken(claims)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return "", err
	}

	log.Println("Token successfully generated")
	return token, nil
}

// GenerateToken generates an empty token.
func (g *GenerateURLController) GenerateToken() (string, error) {
	log.Println("Generating empty token...")
	return utils.GenerateEmptyToken()
}

// createClaims constructs JWT claims.
func (g *GenerateURLController) createClaims(data string) dtos.Claims {
	return dtos.Claims{
		Data:             data,
		RegisteredClaims: utils.GenerateRegisteredClaims("your-app-name", time.Hour),
	}
}

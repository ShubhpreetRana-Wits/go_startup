package repository

import (
	"fmt"
	"log"

	"example.com/startup/internal/domain/entities"
	"example.com/startup/internal/domain/repositories"
	"example.com/startup/internal/infrastucture/db"
	"gorm.io/gorm"
)

type GenerateURLRepositoryImpl struct {
	dbClient *db.PostgresClient
}

// DeleteUrlRequest implements repositories.GenerateURLRepository.
func (g *GenerateURLRepositoryImpl) DeleteUrlRequest(requestId string) error {
	// Perform the delete operation on the database using the requestId
	result := g.dbClient.DB.Where("id = ?", requestId).Delete(&entities.GeneratedUrl{})

	// Check for errors in the delete operation
	if result.Error != nil {
		log.Printf("Error deleting URL request with ID %s: %v", requestId, result.Error)
		return fmt.Errorf("could not delete URL request: %v", result.Error)
	}

	// Check if any rows were deleted
	if result.RowsAffected == 0 {
		log.Printf("No URL request found with ID %s", requestId)
		return fmt.Errorf("no URL request found with ID %s", requestId)
	}

	log.Printf("Successfully deleted URL request with ID %s", requestId)
	return nil
}

// GetUrlRequest implements repositories.GenerateURLRepository.
func (g *GenerateURLRepositoryImpl) GetUrlRequest(requestId string) (*entities.GeneratedUrl, error) {
	var urlRequest entities.GeneratedUrl
	// Perform the fetch operation from the database using the requestId
	if err := g.dbClient.DB.Where("id = ?", requestId).First(&urlRequest).Error; err != nil {
		log.Printf("Error retrieving URL request with ID %s: %v", requestId, err)
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("URL request not found")
		}
		return nil, fmt.Errorf("could not retrieve URL request: %v", err)
	}
	log.Printf("Successfully retrieved URL request with ID %s", requestId)
	return &urlRequest, nil
}

// SaveUrlRequest implements repositories.GenerateURLRepository.
func (g *GenerateURLRepositoryImpl) SaveUrlRequest(requestData *entities.GeneratedUrl) (*entities.GeneratedUrl, error) {
	// Perform the save operation on the database
	if err := g.dbClient.DB.Create(&requestData).Error; err != nil {
		log.Printf("Error saving URL request: %v", err)
		return nil, fmt.Errorf("could not save URL request: %v", err)
	}
	log.Printf("Successfully saved URL request with ID %s", requestData.ID)
	return requestData, nil
}

// NewGenerateUrlRepository creates a new instance of GenerateURLRepositoryImpl.
func NewGenerateUrlRepository(dbClient *db.PostgresClient) repositories.GenerateURLRepository {
	return &GenerateURLRepositoryImpl{
		dbClient: dbClient,
	}
}

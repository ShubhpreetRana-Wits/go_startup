package usecases

import (
	"log"

	"example.com/startup/internal/domain/entities"
	"example.com/startup/internal/domain/repositories"
)

type GeneratedUrlUseCase interface {
	SaveUrlRequest(requestData *entities.GeneratedUrl) (*entities.GeneratedUrl, error)
	GetUrlRequest(requestId string) (*entities.GeneratedUrl, error)
	DeleteUrlRequest(requestId string) error
}

type GeneratedUrlUseCaseImpl struct {
	genrateUrlRepo repositories.GenerateURLRepository
}

func NewGenrateUrlUsecase(genrateUrlRepo repositories.GenerateURLRepository) GeneratedUrlUseCase {
	return &GeneratedUrlUseCaseImpl{
		genrateUrlRepo: genrateUrlRepo,
	}
}

// DeleteUrlRequest implements GeneratedUrlUseCase.
func (g *GeneratedUrlUseCaseImpl) DeleteUrlRequest(requestId string) error {
	err := g.genrateUrlRepo.DeleteUrlRequest(requestId)
	if err != nil {
		log.Printf("Error retrieving from database: %v", err)
		return err
	}

	log.Println("URL data successfully Deleted from storage")
	return nil
}

// GetUrlRequest implements GeneratedUrlUseCase.
func (g *GeneratedUrlUseCaseImpl) GetUrlRequest(requestId string) (*entities.GeneratedUrl, error) {
	response, err := g.genrateUrlRepo.GetUrlRequest(requestId)
	if err != nil {
		log.Printf("Error retrieving from database: %v", err)
		return nil, err
	}

	log.Println("URL data successfully retrieved from storage")
	return response, nil

}

// SaveUrlRequest implements GeneratedUrlUseCase.
func (g *GeneratedUrlUseCaseImpl) SaveUrlRequest(requestData *entities.GeneratedUrl) (*entities.GeneratedUrl, error) {

	response, err := g.genrateUrlRepo.SaveUrlRequest(requestData)
	if err != nil {
		log.Printf("Error saving to database: %v", err)
		return nil, err
	}

	log.Println("URL data successfully saved to storage")
	return response, nil
}

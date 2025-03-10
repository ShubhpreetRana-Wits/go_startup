package repositories

import (
	"example.com/startup/internal/domain/entities"
)

type GenerateURLRepository interface {
	SaveUrlRequest(requestData *entities.GeneratedUrl) (*entities.GeneratedUrl, error)
	GetUrlRequest(requestId string) (*entities.GeneratedUrl, error)
	DeleteUrlRequest(requestId string) error
}

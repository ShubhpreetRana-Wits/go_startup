package entities

import (
	"example.com/startup/internal/dtos"
)

type GeneratedUrl struct {
	BaseModel
	UserID      string  `gorm:"not null"`
	RequestType string  `gorm:"not null"`
	Error       *string `gorm:"null"` // JSON Representation of the error
}

func (GeneratedUrl) TableName() string {
	return "generated_dynamic_url"
}

// ToAuditLogResponseDTO converts the AuditLog model into a response DTO
func (o *GeneratedUrl) ToResponseDTO(token string, url string) dtos.GenerateUrlResponse {
	return dtos.GenerateUrlResponse{
		RedirectURL: url + "?token=" + token,
	}
}

// ToAuditLogResponseDTO converts the AuditLog model into a response DTO
func (o *GeneratedUrl) ToResponseInfoDTO() dtos.GenerateUrlInfoResponse {
	return dtos.GenerateUrlInfoResponse{
		UserID:      o.UserID,
		RequestType: o.RequestType,
	}
}

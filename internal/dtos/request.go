package dtos

type GenerateUrlRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	RequestType string `json:"request_type" validate:"required"`
}

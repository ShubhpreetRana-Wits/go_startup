package dtos

// BaseResponse is the common response structure.
type BaseResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code,omitempty"` // Optional error code
	Data    any    `json:"data,omitempty"` // Optional data
}

// GenerateUrlResponse extends BaseResponse for URL generation.
type GenerateUrlResponse struct {
	RedirectURL string `json:"redirectUrl,omitempty"`
}

// GenerateUrlInfoResponse holds user-specific URL information.
type GenerateUrlInfoResponse struct {
	UserID      string `json:"user_id"`
	RequestType string `json:"request_type"`
}

// ErrorResponse extends BaseResponse for error cases.
type ErrorResponse struct {
	ErrorDetails string `json:"error_details,omitempty"`
}

type FieldErrorResponseDTO struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMsg"`
	Field        string `json:"field"`
}

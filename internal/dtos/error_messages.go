package dtos

import "fmt"

// GetErrorMessage returns the error message related to a specific HTTP status code.
func GetErrorMessage(code int) string {
	statusMessages := map[int]string{
		// 4XX Client Errors
		400: "Bad Request: Invalid request format.",
		401: "Unauthorized: Authentication required.",
		403: "Forbidden: Access is denied.",
		404: "Not Found: Resource does not exist.",
		405: "Method Not Allowed: Unsupported request method.",
		409: "Conflict: Duplicate or conflicting request.",
		422: "Unprocessable Entity: Validation failed.",

		// 5XX Server Errors
		500: "Internal Server Error: Something went wrong.",
		502: "Bad Gateway: Invalid upstream response.",
		503: "Service Unavailable: Server is overloaded or under maintenance.",
		504: "Gateway Timeout: No response from upstream service.",
	}

	if message, exists := statusMessages[code]; exists {
		return message
	}
	return fmt.Sprintf("Something went wrong%d", code)
}

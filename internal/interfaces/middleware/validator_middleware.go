package middleware

import (
	"fmt"

	"example.com/startup/internal/dtos"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateParams[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto T

		// Check if contextData is available
		if contextData := c.Locals("contextData"); contextData != nil {
			if castedDto, ok := contextData.(*T); ok {
				dto = *castedDto
			}
		}

		if err := c.BodyParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"fieldErrors": []dtos.FieldErrorResponseDTO{
					{
						Field:        "requestBody",
						ErrorMessage: "Invalid request body",
						ErrorCode:    "InvalidRequestBody",
					},
				},
			})
		}

		// Extract query parameters
		if err := c.QueryParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"fieldErrors": []dtos.FieldErrorResponseDTO{
					{
						Field:        "requestParams",
						ErrorMessage: "Failed to read request params",
						ErrorCode:    "InvalidRequestParams",
					},
				},
			})
		}

		// Extract path parameters
		if err := c.ParamsParser(&dto); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"fieldErrors": []dtos.FieldErrorResponseDTO{
					{
						Field:        "requestParams",
						ErrorMessage: "Failed to read request params",
						ErrorCode:    "InvalidRequestParams",
					},
				},
			})
		}
		var validate = validator.New()
		// Validate the DTO
		if err := validate.Struct(dto); err != nil {
			var errorMessages []dtos.FieldErrorResponseDTO

			for _, err := range err.(validator.ValidationErrors) {
				errorMessages = append(errorMessages, dtos.FieldErrorResponseDTO{
					Field:        err.Field(),
					ErrorMessage: fmt.Sprintf("Validation failed: '%s'", err.Tag()),
				})
			}

			if len(errorMessages) > 0 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"fieldErrors": errorMessages,
				})
			}
		}

		// Store the validated DTO in the context
		c.Locals("contextData", &dto)

		// Continue to the next handler
		return c.Next()
	}
}

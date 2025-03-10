package middleware

import (
	"log"

	"example.com/startup/internal/dtos"
	"example.com/startup/utils"
	"github.com/gofiber/fiber/v2"
)

// JWTMiddleware is the middleware to validate the JWT token
func JWTUrlMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Queries()["token"]
		if tokenString == "" {
			log.Println("Missing token in query parameters")
			return dtos.NewResponse(c, "Token is required", fiber.StatusBadRequest, "")
		}

		// Validate token and extract claims
		if valid := utils.ValidateToken(tokenString); !valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Extract claims
		claims, err := utils.ExtractClaims(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Failed to extract claims",
			})
		}
		// Store the claims in the locals so they can be used later in the handler
		c.Locals("metadata", claims["data"].(string))

		// Continue to the next handler
		return c.Next()
	}
}

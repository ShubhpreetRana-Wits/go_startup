package middleware

import (
	"log"
	"strings"

	"example.com/startup/internal/dtos"
	"example.com/startup/utils"
	"github.com/gofiber/fiber/v2"
)

// JWTMiddleware is the middleware to validate the JWT token
func JWTMetadataMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		encryptedMetadata, ok := c.Locals("metadata").(string)
		if !ok {
			log.Println("Metadata not found in token")
			return dtos.NewResponse(c, "metadata not found in token", fiber.StatusBadRequest, "")
		}

		// Clean null characters if needed

		decryptedData, err := utils.Decrypt(encryptedMetadata, utils.JWTSecret)
		if err != nil {
			log.Printf("Error decrypting metadata: %v", err)
			return dtos.NewResponse(c, err.Error(), fiber.StatusBadRequest, "")
		}
		decryptedData = strings.ReplaceAll(decryptedData, "\x00", "")

		log.Println("JWT successfully extracted and decrypted")
		// Store the claims in the locals so they can be used later in the handler
		c.Locals("jwtdata", decryptedData)

		return c.Next()
	}
}

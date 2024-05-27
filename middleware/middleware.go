package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-authentication/utils"
)

func Protected(c *fiber.Ctx) error {

	// token := "c"

	token := c.Cookies("X-Auth-Token")

	_, err := utils.VerifyToken(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Invalid token",
		})
	}

	return c.Next()

}

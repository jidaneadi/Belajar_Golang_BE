package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jidaneadi/projectkp-backend/utils"
)

func Auth(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	// Memeriksa keberadaan header Authorization
	if authorization == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "Missing Authorization header"})
	}

	// Memisahkan Bearer dan token dari header Authorization
	tokenString := ""
	splitToken := strings.Split(authorization, "Bearer ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	}

	// Memeriksa keberadaan token
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "Missing token"})
	}

	verify, err := utils.VerifyAccesToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg": "Invalid Acces Token",
		})
	}

	c.Locals("jwt", verify.Claims)
	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {
	return c.Next()
}

package middlewares

import (
	"maytri/models"
	"strings"

	"github.com/MelloB1989/karma/config"
	"github.com/MelloB1989/karma/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// JWTMiddleware is the middleware function to verify JWT tokens
func IsUserVerified(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
			"error":   errors.NewKarmaError().GetError(4002),
		})
	}

	// Extract the token from the Bearer string
	tokenStr := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
			"error":   errors.NewKarmaError().GetError(4002),
		})
	}

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&models.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			// Make sure the token's algorithm is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(config.DefaultConfig().JWTSecret), nil
		},
	)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
			"error":   errors.NewKarmaError().GetError(4002),
		})
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		// Store the claims in the context's locals
		// fmt.Println(claims.UserID)
		c.Locals("uid", claims.Uid)
		c.Locals("email", claims.Email)
		c.Locals("age", claims.Age)
		c.Locals("phone", claims.Phone)
		c.Locals("name", claims.Name)
		// Continue with the next handler
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
		"error":   errors.NewKarmaError().GetError(4002),
	})
}

package routes

import (
	"maytri/internal/handlers/users"
	"maytri/internal/helpers/auth"
	"maytri/middlewares"
	"maytri/models"

	a "github.com/MelloB1989/karma/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, X-Karma-Admin-Auth",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(models.ResponseHTTP{
			Success: true,
			Data:    nil,
			Message: "OK",
		})
	})
	v1 := app.Group("/v1")

	authRoutes := v1.Group("/auth")
	authRoutes.Post("/login", a.LoginWithPhoneOTPHandler(auth.GetUser))
	authRoutes.Post("/verify_otp", a.VerifyPhoneOTPHandler(auth.GetUser))
	authRoutes.Post("/register", middlewares.IsUserVerified, users.RegisterUser)

	return app
}

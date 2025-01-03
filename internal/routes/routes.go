package routes

import (
	"maytri/internal/handlers/ai"
	"maytri/internal/handlers/chat"
	"maytri/internal/handlers/maytri"
	"maytri/internal/handlers/messages"
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

	chatRoutes := v1.Group("/chat")
	chatRoutes.Post("/private", middlewares.IsUserVerified, chat.CreatePrivateChat)

	aiRoutes := v1.Group("/ai")
	aiRoutes.Post("/newMaytri", middlewares.IsUserVerified, ai.GetNewMaytri)

	messageRoutes := v1.Group("/message")
	messageRoutes.Post("/private/:chat_id", middlewares.IsUserVerified, messages.PrivateChat)
	messageRoutes.Get("/private/:chat_id", middlewares.IsUserVerified, messages.GetPrivateChatMessages)

	maytriRoutes := v1.Group("/maytri")
	maytriRoutes.Get("/", middlewares.IsUserVerified, maytri.GetMyMaytris)
	return app
}

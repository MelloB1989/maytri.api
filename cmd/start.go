package cmd

import (
	"maytri/internal/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer() {
	app := routes.Routes()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Listen(":9000")
}

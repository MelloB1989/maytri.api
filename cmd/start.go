package cmd

import (
	"maytri/internal/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := routes.Routes()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Listen(":3000")
}

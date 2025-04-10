package routes

import (
	"go-fiber-vercel/handlers"

	"github.com/gofiber/fiber/v2"
)

// V1Route mengatur routing untuk versi 1
func V1Route(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("/", handlers.V1Handler)
}

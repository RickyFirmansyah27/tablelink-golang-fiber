package routes

import (
	"go-fiber-vercel/handlers"

	"github.com/gofiber/fiber/v2"
)

// V2Route mengatur routing untuk versi 2
func V2Route(app *fiber.App) {
	v2 := app.Group("/v2")
	v2.Get("/", handlers.V2Handler)
}

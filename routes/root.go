package routes

import (
	"go-fiber-vercel/handlers"

	"github.com/gofiber/fiber/v2"
)

func RootRoute(app *fiber.App) {
	app.Get("/", handlers.RootHandler)
}

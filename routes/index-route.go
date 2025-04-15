package routes

import (
	"go-fiber-vercel/handlers"

	"github.com/gofiber/fiber/v2"
)

func RootRoute(app *fiber.App) {
	app.Get("/", handlers.RootHandler)
	app.Get("/items", handlers.GetItems)
	app.Post("/items", handlers.CreateItem)
	app.Put("/items", handlers.UpdateItem)
	app.Delete("/items/:uuid", handlers.DeleteItem)

	// Ingredient
	app.Get("/ingredient", handlers.GetIngredient)
	app.Post("/ingredient", handlers.CreateIngredient)
	app.Put("/ingredient", handlers.UpdateIngredient)
	app.Delete("/ingredient/:uuid", handlers.DeleteIngredient)
}

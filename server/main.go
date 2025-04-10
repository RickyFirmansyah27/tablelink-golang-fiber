package main

import (
	"go-fiber-vercel/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	routes.RootRoute(app)
	routes.V1Route(app)
	routes.V2Route(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Fiber server started on %s...", port)
	log.Fatal(app.Listen(":" + port))
}

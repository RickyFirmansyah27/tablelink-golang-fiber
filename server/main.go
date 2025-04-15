package main

import (
	"go-fiber-vercel/routes"
	"go-fiber-vercel/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := config.DBConnection()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	app := fiber.New()

	routes.RootRoute(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Fiber server started on %s...", port)
	log.Fatal(app.Listen(":" + port))
}

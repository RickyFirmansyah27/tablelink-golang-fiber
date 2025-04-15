package handler

import (
	"go-fiber-vercel/routes"
	"net/http"
	"log"
	"go-fiber-vercel/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := fiber.New()

	err := config.DBConnection()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	routes.RootRoute(app)
	// add more version

	return adaptor.FiberApp(app)
}

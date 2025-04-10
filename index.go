package handler

import (
	"go-fiber-vercel/routes"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := fiber.New()

	routes.RootRoute(app)
	routes.V1Route(app)
	routes.V2Route(app)
	// add more version

	return adaptor.FiberApp(app)
}

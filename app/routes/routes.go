package routes

import (
	"golang-vercel/app/controller"
	"golang-vercel/app/helpers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	app.NoRoute(NoRouteAccess)

	app.GET("/", func(ctx *gin.Context) {
        helpers.Success(ctx, "Welcome to Go Vercel", nil)
    })

	app.GET("/ping", controller.Ping)

	route := app.Group("/api")
	{
		route.GET("/hello/:name", controller.Hello)
	}
}

func NoRouteAccess(c *gin.Context) {
	helpers.NotFound(c, "No Route Found")
}

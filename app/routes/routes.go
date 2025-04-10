package routes

import (
	"golang-vercel/app/controller"
	"golang-vercel/app/helpers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	app.NoRoute(ErrRouter)

	app.GET("/", func(ctx *gin.Context) {
        helpers.Success(ctx, "Welcome to Go Vercel", nil)
    })

	app.GET("/ping", controller.Ping)

	route := app.Group("/api")
	{
		route.GET("/hello/:name", controller.Hello)
	}
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}

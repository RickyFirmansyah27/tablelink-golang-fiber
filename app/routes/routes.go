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

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

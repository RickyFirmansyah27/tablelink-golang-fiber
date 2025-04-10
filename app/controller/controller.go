package controller

import (
	"net/http"

	"golang-vercel/app/helpers"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	helpers.Success(ctx, "Ping Pong", nil)
}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello %v", c.Param("name"))
}

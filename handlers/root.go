package handlers

import "github.com/gofiber/fiber/v2"
import "go-fiber-vercel/helpers"

// RootHandler mengatur rute root
func RootHandler(ctx *fiber.Ctx) error {
	return helpers.Success(ctx, "Welcome to Go Fiber Vercel", nil)
}

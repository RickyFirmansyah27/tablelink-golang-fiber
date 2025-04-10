package handlers

import "github.com/gofiber/fiber/v2"
import "go-fiber-vercel/helpers"

// V1Handler mengatur rute untuk versi 1
func V1Handler(ctx *fiber.Ctx) error {
	return helpers.Success(ctx, "Version 1 Route", nil)
}

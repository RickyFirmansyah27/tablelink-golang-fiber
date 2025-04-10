package handlers

import "github.com/gofiber/fiber/v2"
import "go-fiber-vercel/helpers"

// V2Handler mengatur rute untuk versi 2
func V2Handler(ctx *fiber.Ctx) error {
	return helpers.Success(ctx, "Version 1 Route", nil)
}

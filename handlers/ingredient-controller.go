package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-vercel/service"
	"go-fiber-vercel/helpers"
	"go-fiber-vercel/models"
	"log"
)

func GetIngredient(c *fiber.Ctx) error {
	log.Printf("[ItemsController] - Incoming request with query params: %v", c.Queries())

	totalData, items, err := service.GetIngredient(c)
	if err != nil {
		log.Printf("[ItemsController] - Failed to fetch items: %v", err)
		return helpers.Error(c, 400, "Failed to fetch items", err)
	}

	data := []any{
		fiber.Map{
			"total_data": totalData,
			"items":      items,
		},
	}

	log.Printf("[ItemsController] - Successfully fetched items")

	return helpers.Success(c, "Successfully fetched items", data)
}

func CreateIngredient(c *fiber.Ctx) error {
	var input models.Ingredient
	if err := c.BodyParser(&input); err != nil {
		log.Printf("[ItemsController] - Error parsing request body: %v", err)
		return helpers.Error(c, 400, "Invalid request body", err)
	}

	item, err := service.CreateIngredient(input)
	if err != nil {
		log.Printf("[ItemsController] - Failed to create item: %v", err)
		return helpers.Error(c, 500, "Failed to create item", err)
	}

	return helpers.Success(c, "Item successfully created", item)
}

func UpdateIngredient(c *fiber.Ctx) error {
	var input models.Ingredient
	if err := c.BodyParser(&input); err != nil {
		log.Printf("[ItemsController] - Error parsing request body: %v", err)
		return helpers.Error(c, 400, "Invalid request body", err)
	}

	item, err := service.UpdateIngredient(input)
	if err != nil {
		log.Printf("[ItemsController] - Failed to udpated item: %v", err)
		return helpers.Error(c, 500, "Failed to udpated item", err)
	}

	return helpers.Success(c, "Item successfully udpated", item)
}

func DeleteIngredient(c *fiber.Ctx) error {
    uuid := c.Params("uuid")
    if uuid == "" {
        return helpers.Error(c, 400, "UUID is required", nil)
    }

    err := service.DeleteIngredientByUUID(uuid)
    if err != nil {
        log.Printf("[ItemsController] - Error deleting item: %v", err)
        return helpers.Error(c, 500, "Failed to delete item", err)
    }

    log.Printf("[ItemsController] - Successfully deleted item with UUID: %s", uuid)
    return helpers.Success(c, "Item successfully deleted", nil)
}

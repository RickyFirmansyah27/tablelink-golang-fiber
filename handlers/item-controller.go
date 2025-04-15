package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-vercel/service"
	"go-fiber-vercel/helpers"
	"go-fiber-vercel/models"
	"log"
)

func RootHandler(ctx *fiber.Ctx) error {
	return helpers.Success(ctx, "Welcome to Go Fiber Vercel", nil)
}

func GetItems(c *fiber.Ctx) error {
	log.Printf("[ItemsController] - Incoming request with query params: %v", c.Queries())

	totalData, items, err := service.GetItems(c)
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

func CreateItem(c *fiber.Ctx) error {
	var input models.Item
	if err := c.BodyParser(&input); err != nil {
		log.Printf("[ItemsController] - Error parsing request body: %v", err)
		return helpers.Error(c, 400, "Invalid request body", err)
	}

	item, err := service.CreateItem(input)
	if err != nil {
		log.Printf("[ItemsController] - Failed to create item: %v", err)
		return helpers.Error(c, 500, "Failed to create item", err)
	}

	return helpers.Success(c, "Item successfully created", item)
}

func UpdateItem(c *fiber.Ctx) error {
    uuid := c.Params("uuid")
    if uuid == "" {
        return helpers.Error(c, 400, "UUID is required", nil)
    }

    var input struct {
        Name   string `json:"name"`
        Price  int    `json:"price"`
        Status string `json:"status"`
    }

    itemInput := models.Item{
        Name:   input.Name,
        Price:  input.Price,
        Status: input.Status,
    }

    err := service.UpdateItem(uuid, itemInput)
    if err != nil {
        log.Printf("[ItemsController] - Error updating item: %v", err)
        return helpers.Error(c, 500, "Failed to update item", err)
    }

    log.Printf("[ItemsController] - Successfully updated item with UUID: %s", uuid)
    return helpers.Success(c, "Item successfully updated", itemInput)
}

func DeleteItem(c *fiber.Ctx) error {
    uuid := c.Params("uuid")
    if uuid == "" {
        return helpers.Error(c, 400, "UUID is required", nil)
    }

    err := service.DeleteItemByUUID(uuid)
    if err != nil {
        log.Printf("[ItemsController] - Error deleting item: %v", err)
        return helpers.Error(c, 500, "Failed to delete item", err)
    }

    log.Printf("[ItemsController] - Successfully deleted item with UUID: %s", uuid)
    return helpers.Success(c, "Item successfully deleted", nil)
}

package service

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-vercel/config"
	"go-fiber-vercel/models"
	"github.com/google/uuid"
	"log"
)

func GetIngredient(c *fiber.Ctx) (int, []models.Ingredient, error) {
	log.Println("[ItemsService] - Fetching items...", c.Queries())

	itemsMap, totalData, err := config.GetAllIngredient(c)
	if err != nil {
		log.Printf("[ItemsService] - Error fetching items: %v", err)
		return 0, nil, err
	}

	items := make([]models.Ingredient, 0, len(itemsMap))
	for _, itemMap := range itemsMap {
		item := models.Ingredient{
			UUID:        itemMap["uuid"].(string),
			Name:        itemMap["name"].(string),
			Cause_Alergy: itemMap["cause_alergy"].(bool),
			Type:        itemMap["type"].(string),
			Status:      itemMap["status"].(string),
		}
		items = append(items, item)
	}

	log.Printf("[ItemsService] - Successfully fetched %d items", len(items))
	return totalData, items, nil
}

func CreateIngredient(input models.Ingredient) (models.Ingredient, error) {
	log.Printf("[ItemsService] - Creating new item with Name: %s", input.Name)

	itemUUID := uuid.New()

	item := models.Ingredient{
		UUID:   itemUUID.String(),
		Name:   input.Name,
		Type: input.Type,
		Status: input.Status,
	}

	err := config.CreateIngredient(item)

	if err != nil {
		log.Printf("[ItemsService] - Error creating item: %v", err)
		return models.Ingredient{}, err
	}

	log.Printf("[ItemsService] - Successfully created new item with UUID: %s", item.UUID)

	return item, nil
}


func UpdateIngredient(item models.Ingredient) (*models.Ingredient, error) {
	log.Printf("[ItemsService] - Updating item with UUID: %s", item.UUID)

	err := config.UpdateIngredientByUUID(item)
	if err != nil {
		log.Printf("[ItemsService] - Failed to update item: %s, error: %v", err)
		return nil, err
	}

	log.Printf("[ItemsService] - Successfully updated item")
	return &item, nil
}

func DeleteIngredientByUUID(uuid string) error {
	log.Printf("[ItemsService] - Deleting item with UUID: %s", uuid)

	err := config.DeleteIngredientByUUID(uuid)
	if err != nil {
		log.Printf("[ItemsService] - Failed to Deleting item with UUID: %s, error: %v", uuid, err)
		return err
	}

	log.Printf("[ItemsService] - Successfully Deleting item with UUID: %s", uuid)
	return nil
}

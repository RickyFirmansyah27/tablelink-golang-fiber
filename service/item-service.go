package service

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-vercel/config"
	"go-fiber-vercel/models"
	"github.com/google/uuid"
	"log"
)

func GetItems(c *fiber.Ctx) (int, []models.Item, error) {
	log.Println("[ItemsService] - Fetching items...", c.Queries())

	itemsMap, totalData, err := config.GetAllItems(c)
	if err != nil {
		log.Printf("[ItemsService] - Error fetching items: %v", err)
		return 0, nil, err
	}

	items := make([]models.Item, 0, len(itemsMap))
	for _, itemMap := range itemsMap {
		item := models.Item{
			UUID:   itemMap["uuid"].(string),
			Name:   itemMap["name"].(string),
			Price:  itemMap["name"].(int),
			Status:  itemMap["name"].(string),
		}
		items = append(items, item)
	}

	log.Printf("[ItemsService] - Successfully fetched %d items", len(items))
	return totalData, items, nil
}

func getStringValue(itemMap fiber.Map, key string) string {
	if val, ok := itemMap[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func CreateItem(input models.Item) (models.Item, error) {
	log.Printf("[ItemsService] - Creating new item with Name: %s", input.Name)

	itemUUID := uuid.New()

	item := models.Item{
		UUID:   itemUUID.String(),
		Name:   input.Name,
		Price:  input.Price,
		Status: input.Status,
	}

	err := config.CreateItem(item)

	if err != nil {
		log.Printf("[ItemsService] - Error creating item: %v", err)
		return models.Item{}, err
	}

	log.Printf("[ItemsService] - Successfully created new item with UUID: %s", item.UUID)

	return item, nil
}


func UpdateItem(item models.Item) (*models.Item, error) {
	log.Printf("[ItemsService] - Updating item with UUID: %s", item.UUID)

	err := config.UpdateItemByUUID(item)
	if err != nil {
		log.Printf("[ItemsService] - Failed to update item: %s, error: %v", err)
		return nil, err
	}

	log.Printf("[ItemsService] - Successfully updated item")
	return &item, nil
}

func DeleteItemByUUID(uuid string) error {
	log.Printf("[ItemsService] - Deleting item with UUID: %s", uuid)

	err := config.DeleteItemByUUID(uuid)
	if err != nil {
		log.Printf("[ItemsService] - Failed to Deleting item with UUID: %s, error: %v", uuid, err)
		return err
	}

	log.Printf("[ItemsService] - Successfully Deleting item with UUID: %s", uuid)
	return nil
}

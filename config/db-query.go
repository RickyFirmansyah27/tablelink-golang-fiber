package config

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"log"

	"github.com/gofiber/fiber/v2"
	"go-fiber-vercel/models"
)

var db *sql.DB

func GetAllItems(c *fiber.Ctx) ([]fiber.Map, int, error) {
	params := c.Queries()

	page, err := strconv.Atoi(params["page"])
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(params["size"])
	if err != nil {
		size = 10
	}

	if size != 10 && size != 20 && size != 50 {
		size = 10
	}

	query := "SELECT uuid, name, price, status FROM tm_item"
	queryCount := "SELECT COUNT(*) FROM tm_item"

	whereClauses := []string{}
	queryParams := []interface{}{}
	paramIndex := 1

	if name, ok := params["name"]; ok && name != "" {
		whereClauses = append(whereClauses, "name ILIKE $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, "%"+name+"%")
		paramIndex++
	}

	if price, ok := params["price"]; ok && price != "" {
		whereClauses = append(whereClauses, "CAST(price AS TEXT) ILIKE $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, "%"+price+"%")
		paramIndex++
	}

	if status, ok := params["status"]; ok && status != "" {
		whereClauses = append(whereClauses, "CAST(status AS TEXT) ILIKE $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, "%"+status+"%")
		paramIndex++
	}

	if len(whereClauses) > 0 {
		whereStr := strings.Join(whereClauses, " AND ")
		query += " WHERE " + whereStr
		queryCount += " WHERE " + whereStr
	}

	sortBy := params["sort_by"]
	sortOrder := params["sort_order"]

	allowedSortFields := map[string]bool{
		"uuid": true, "name": true, "price": true, "status": true,
	}

	if _, ok := allowedSortFields[sortBy]; !ok {
		sortBy = "uuid"
	}

	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "ASC"
	}

	query += " ORDER BY " + sortBy + " " + sortOrder

	offset := (page - 1) * size
	query += " LIMIT $" + strconv.Itoa(paramIndex) + " OFFSET $" + strconv.Itoa(paramIndex+1)
	queryParams = append(queryParams, size, offset)

	countRows, err := ExecuteSQLWithParams(queryCount, queryParams[:paramIndex-1]...)
	if err != nil {
		return nil, 0, err
	}
	defer countRows.Close()

	var totalData int
	if countRows.Next() {
		if err := countRows.Scan(&totalData); err != nil {
			return nil, 0, err
		}
	}

	rows, err := ExecuteSQLWithParams(query, queryParams...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := []fiber.Map{}
	for rows.Next() {
		var uuid, name string
		var price float64
		var status int

		if err := rows.Scan(&uuid, &name, &price, &status); err != nil {
			return nil, 0, err
		}

		items = append(items, fiber.Map{
			"uuid":   uuid,
			"name":   name,
			"price":  fmt.Sprintf("%.2f", price),
			"status": strconv.Itoa(status),
		})
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return items, totalData, nil
}

func CreateItem(item models.Item) error {
	query := `
		INSERT INTO tm_item (uuid, name, price, status) 
		VALUES ($1, $2, $3, $4)
	`

	queryParams := []interface{}{
		item.UUID,
		item.Name,
		item.Price,
		item.Status,
	}

	_, err := ExecuteSQLWithParams(query, queryParams...)
	if err != nil {
		log.Printf("[CreateItem] - Error executing query: %v", err)
		return err
	}

	log.Printf("[CreateItem] - Successfully created item with UUID: %s", item.UUID,)
	return nil
}


func UpdateItemByUUID(item models.Item) error {
	query := `UPDATE tm_item 
		SET name = $1, price = $2, status = $3
		WHERE uuid = $4`

	queryParams := []interface{}{
		item.Name,
		item.Price,
		item.Status,
		item.UUID,
	}

	_, err := ExecuteSQLWithParams(query, queryParams...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteItemByUUID(uuid string) error {
	query := `Delete from tm_item WHERE uuid = $1`

	_, err := ExecuteSQLWithParams(query, uuid)
	if err != nil {
		log.Printf("[Config] - Error soft deleting item with UUID: %s, error: %v", uuid, err)
		return err
	}

	log.Printf("[Config] - Successfully soft deleted item with UUID: %s", uuid)
	return nil
}
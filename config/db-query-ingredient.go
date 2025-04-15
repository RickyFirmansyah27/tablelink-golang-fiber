package config

import (
	"database/sql"
	"strconv"
	"strings"
	"log"

	"github.com/gofiber/fiber/v2"
	"go-fiber-vercel/models"
)

var dbms *sql.DB

func GetAllIngredient(c *fiber.Ctx) ([]fiber.Map, int, error) {
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

	query := "SELECT uuid, name, cause_alergy, type, status FROM tm_ingredient"
	queryCount := "SELECT COUNT(*) FROM tm_ingredient"

	whereClauses := []string{}
	queryParams := []interface{}{}
	paramIndex := 1

	if name, ok := params["name"]; ok && name != "" {
		whereClauses = append(whereClauses, "name ILIKE $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, "%"+name+"%")
		paramIndex++
	}

	if cause_alergy, ok := params["cause_alergy"]; ok && cause_alergy != "" {
		whereClauses = append(whereClauses, "CAST(cause_alergy AS TEXT) ILIKE $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, "%"+cause_alergy+"%")
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
		"uuid": true, "name": true, "cause_alergy": true, "type": true, "status": true,
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
		var cause_alergy bool
		var status int
		var types int

		if err := rows.Scan(&uuid, &name, &cause_alergy, &types, &status); err != nil {
			return nil, 0, err
		}

		items = append(items, fiber.Map{
			"uuid":        uuid,
			"name":        name,
			"cause_alergy": cause_alergy,
			"type":        strconv.Itoa(types),
			"status":      strconv.Itoa(status),
		})
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return items, totalData, nil
}

func CreateIngredient(item models.Ingredient) error {
	query := `
		INSERT INTO tm_ingredient (uuid, name, cause_alergy, type, status) 
		VALUES ($1, $2, $3, $4, $5)
	`

	queryParams := []interface{}{
		item.UUID,
		item.Name,
		item.Cause_Alergy,
		item.Type,
		item.Status,
	}

	_, err := ExecuteSQLWithParams(query, queryParams...)
	if err != nil {
		log.Printf("[CreateIngredient] - Error executing query: %v", err)
		return err
	}

	log.Printf("[CreateIngredient] - Successfully created item with UUID: %s", item.UUID)
	return nil
}

func UpdateIngredientByUUID(item models.Ingredient) error {
	query := `UPDATE tm_ingredient 
		SET name = $1, cause_alergy = $2, status = $3, type = $4
		WHERE uuid = $5`

	queryParams := []interface{}{
		item.Name,
		item.Cause_Alergy,
		item.Status,
		item.Type,
		item.UUID,
	}

	_, err := ExecuteSQLWithParams(query, queryParams...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteIngredientByUUID(uuid string) error {
	query := `DELETE FROM tm_ingredient WHERE uuid = $1`

	_, err := ExecuteSQLWithParams(query, uuid)
	if err != nil {
		log.Printf("[Config] - Error soft deleting item with UUID: %s, error: %v", uuid, err)
		return err
	}

	log.Printf("[Config] - Successfully soft deleted item with UUID: %s", uuid)
	return nil
}

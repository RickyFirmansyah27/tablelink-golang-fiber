package models

type Item struct {
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	Price        int `json:"price"`
	Status       string `json:"status"`
}
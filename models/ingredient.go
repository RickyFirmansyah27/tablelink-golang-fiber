package models

type Ingredient struct {
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	Cause_Alergy        bool `json:"cause_alergy"`
	Type       string `json:"type"`
	Status       string `json:"status"`
}
package models

type Beer struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Style       string `json:"style"`
	Description string `json:"description"`
}

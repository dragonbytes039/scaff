package models

type Example struct {
	ID      interface{} `json:"id"`
	Title   string      `json:"title"`
	Content string      `json:"content"`
}

package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

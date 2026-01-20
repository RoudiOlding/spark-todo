package models

import (
	"gorm.io/gorm"
)

// This matches the database table "todos"
type Todo struct {
	gorm.Model        // Adds ID, CreatedAt, UpdatedAt automatically
	Title      string `json:"title"`
	Completed  bool   `json:"completed" gorm:"default:false"`
}

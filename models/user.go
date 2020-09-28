package models

import (
	"gorm.io/gorm"
	"time"
)

// Task - Model of a basic task
type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name string
	Email  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
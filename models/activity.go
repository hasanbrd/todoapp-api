package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `json:"email"`
	Title     string         `gorm:"not null" json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

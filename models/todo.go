package models

import (
	"time"

	"gorm.io/gorm"
)

type ToDo struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	ActivityID uint           `gorm:"not null" json:"activity_group_id"`
	Title      string         `gorm:"not null" json:"title"`
	IsActive   bool           `json:"is_active"`
	Priority   string         `json:"priority"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

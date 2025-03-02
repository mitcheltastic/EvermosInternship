package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"unique;not null"`
	CreatedAt time.Time      
	UpdatedAt time.Time      
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

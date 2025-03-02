package models

import (
	"gorm.io/gorm"
	"time"
)

type Store struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"not null"`
	Name      string         `gorm:"not null"`
	ImageURL  string         
	CreatedAt time.Time      
	UpdatedAt time.Time      
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

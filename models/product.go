package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID            uint           `gorm:"primaryKey"`
	Name          string         `gorm:"not null"`
	Slug          string         `gorm:"unique;not null"`
	ResellerPrice float64        `gorm:"not null"`
	ConsumerPrice float64        `gorm:"not null"`
	Stock         int            `gorm:"not null"`
	Description   string         
	StoreID       uint           `gorm:"not null"`
	CategoryID    uint           `gorm:"not null"`
	Price         float64        
	CreatedAt     time.Time      
	UpdatedAt     time.Time      
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID              uint           `gorm:"primaryKey"`
	UserID          uint           `gorm:"not null"`
	ShippingAddress string         
	TotalPrice      float64        
	InvoiceCode     string         `gorm:"type:varchar(255);not null;unique"`
	PaymentMethod   string         
	Status          string         
	TotalAmount     float64        
	CreatedAt       time.Time      
	UpdatedAt       time.Time      
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

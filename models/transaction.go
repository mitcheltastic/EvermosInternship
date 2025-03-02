package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
    ID              uint           `gorm:"primaryKey"`
    UserID          uint           `gorm:"not null" json:"user_id"`
    ShippingAddress string         `json:"shipping_address"`
    TotalPrice      float64        `json:"total_price"`
    InvoiceCode     string         `gorm:"type:varchar(255);not null;unique" json:"invoice_code"`
    PaymentMethod   string         `json:"payment_method"`
    Status          string         `json:"status"`
    TotalAmount     float64        `json:"total_amount"`
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}


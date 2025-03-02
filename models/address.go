package models

import "gorm.io/gorm"

type Address struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     uint   `json:"user_id"`
	ProvinceID string `json:"province_id"`
	CityID     string `json:"city_id"`
	Detail     string `json:"detail"`
	PostalCode string `json:"postal_code"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `json:"name"`
	Email      string         `gorm:"unique" json:"email"`
	Password   string         `json:"-"`
	Phone      string         `json:"phone"`
	BirthDate  time.Time      `json:"birth_date"`
	Gender     string         `json:"gender"`
	Bio        string         `json:"bio"`
	Job        string         `json:"job"`
	ProvinceID uint           `json:"province_id"`
	CityID     uint           `json:"city_id"`
	IsAdmin    bool           `json:"is_admin"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint           `gorm:"primaryKey"`
	Name       string         `gorm:"not null"`
	Email      string         `gorm:"unique;not null"`
	Password   string         `gorm:"not null"`
	Phone      string         `gorm:"unique;not null"`
	BirthDate  string         
	Gender     string         
	Bio        string         
	Job        string         
	ProvinceID uint           
	CityID     uint           
	IsAdmin    bool           
	CreatedAt  time.Time      
	UpdatedAt  time.Time      
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

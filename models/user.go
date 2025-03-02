package models

import (
	"database/sql/driver"
	"time"
	"gorm.io/gorm"
)

// CustomDate Type to Parse "YYYY-MM-DD"
type CustomDate time.Time

// Convert CustomDate to time.Time
func (cd CustomDate) ToTime() time.Time {
	return time.Time(cd)
}

// Check if CustomDate is empty
func (cd CustomDate) IsZero() bool {
	return cd.ToTime().IsZero()
}

// Format CustomDate as "YYYY-MM-DD"
func (cd CustomDate) Format(layout string) string {
	return cd.ToTime().Format(layout)
}

// Parse JSON "YYYY-MM-DD" into CustomDate
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	parsed, err := time.Parse(`"2006-01-02"`, str)
	if err != nil {
		return err
	}
	*cd = CustomDate(parsed)
	return nil
}

// Convert CustomDate to SQL value for GORM
func (cd CustomDate) Value() (driver.Value, error) {
	return cd.ToTime(), nil
}

// Scan converts SQL value to CustomDate for GORM
func (cd *CustomDate) Scan(value interface{}) error {
	if value == nil {
		*cd = CustomDate(time.Time{})
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*cd = CustomDate(t)
		return nil
	}
	return nil
}

type User struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `json:"name" binding:"required"`
	Email      string         `gorm:"unique" json:"email" binding:"required,email"`
	Password   string         `json:"password" binding:"required"` // 🔥 FIXED: Ensure password is required
	Phone      string         `json:"phone"`
	BirthDate  CustomDate     `json:"birth_date"`
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
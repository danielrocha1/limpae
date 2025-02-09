package models

type Diarist struct {
	ID             uint    `gorm:"primaryKey"`
	UserID         uint    `gorm:"unique;not null"`
	Bio            string
	ExperienceYears int    `gorm:"check:experience_years >= 0"`
	PricePerHour   float64 `gorm:"not null"`
	Available      bool    `gorm:"default:true"`
}

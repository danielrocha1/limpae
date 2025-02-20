package models

type Diarists struct {
	ID             uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"foreignKey:UserID;not null"`
	Bio            string
	ExperienceYears int    `gorm:"check:experience_years >= 0"`
	PricePerHour   float64 `gorm:"not null"`
}

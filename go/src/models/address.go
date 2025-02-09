package models

type Address struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Street    string  `gorm:"size:150;not null"`
	Number    string  `gorm:"size:10"`
	City      string  `gorm:"size:50;not null"`
	State     string  `gorm:"size:2;not null"`
	Zipcode   string  `gorm:"size:10;not null"`
	Latitude  float64 `gorm:"type:decimal(9,6)"`
	Longitude float64 `gorm:"type:decimal(9,6)"`
}

package models

type UserProfile struct {
	ID             uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"foreignKey:UserID;not null"`
	Bio            string
	HouseDescription string 
}

package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"size:100;not null"`
	Photo        string    `gorm:"size:255;unique;not null"`
	Email        string    `gorm:"size:100;unique;not null"`
	Phone        string    `gorm:"size:20;unique;not null"`
	Cpf          string    `gorm:"size:20;unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"size:10;not null;check:role IN ('cliente', 'diarista')"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UserProfile  UserProfile `gorm:"foreignKey:UserID"`
	Diarists     []Diarists    `gorm:"foreignKey:UserID"`
	Address    []Address    `gorm:"foreignKey:UserID"`	
}

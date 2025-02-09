package models

import "time"

type Service struct {
	ID           uint      `gorm:"primaryKey"`
	ClientID     uint      `gorm:"not null"`
	DiaristID    uint      `gorm:"not null"`
	AddressID    uint
	Status       string    `gorm:"size:20;default:'pendente'"`
	TotalPrice   float64   `gorm:"not null"`
	DurationHours float64  `gorm:"not null"`
	ScheduledAt  time.Time `gorm:"not null"`
	CompletedAt  *time.Time
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
}

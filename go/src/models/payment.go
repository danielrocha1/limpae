package models

import "time"

type Payment struct {
	ID         uint      `gorm:"primaryKey"`
	ServiceID  uint      `gorm:"unique;not null"`
	ClientID   uint      `gorm:"not null"`
	DiaristID  uint      `gorm:"not null"`
	Amount     float64   `gorm:"not null"`
	Status     string    `gorm:"size:20;default:'pendente'"`
	Method     string    `gorm:"size:20"`
	PaidAt     *time.Time
}

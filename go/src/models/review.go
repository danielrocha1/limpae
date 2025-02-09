package models

import "time"

type Review struct {
	ID         uint      `gorm:"primaryKey"`
	ServiceID  uint      `gorm:"not null"`
	ReviewerID uint      `gorm:"not null"`
	ReviewedID uint      `gorm:"not null"`
	Rating     int       `gorm:"check:rating BETWEEN 1 AND 5"`
	Comment    string
	CreatedAt  time.Time `gorm:"default:current_timestamp"`
}

package models

import "time"

// Subscription representa um plano de assinatura
type Subscription struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Plan      string    `json:"plan" gorm:"type:varchar(20);not null;check:plan IN ('free', 'basic', 'premium')"`
	Price     float64   `json:"price" gorm:"not null"`
	Status    string    `json:"status" gorm:"type:varchar(20);not null;check:status IN ('active', 'inactive', 'canceled')"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

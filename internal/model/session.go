package model

import "time"

type Session struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	BelongTo  uint      `json:"belong_to"`
}

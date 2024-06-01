package models

import "time"

type Product struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

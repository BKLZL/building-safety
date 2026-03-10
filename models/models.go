package models

import "time"

type Customer struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Industry   string    `json:"industry"`
	Contact    string    `json:"contact"`
	Phone      string    `json:"phone"`
	CreditCode string    `json:"credit_code"`
	Address    string    `json:"address"`
	CreatedAt  time.Time `json:"created_at"`
}

type CheckItem struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Standard    string `json:"standard"`
	IsPhoto     bool   `json:"is_photo"`
	IsVideo     bool   `json:"is_video"`
}

package model

import (
	"time"
)

type Entry struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	User      string    `json:"user"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

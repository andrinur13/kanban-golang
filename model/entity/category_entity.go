package entity

import (
	"time"
)

type Category struct {
	ID        int       `json:"id"`
	Type      string    `json:"type" valid:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

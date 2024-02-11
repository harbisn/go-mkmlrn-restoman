package models

import (
	"time"
)

type Menu struct {
	ID          uint64    `json:"id" pg:"id, pk"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Price       int32     `json:"price"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedBy   string    `json:"UpdatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

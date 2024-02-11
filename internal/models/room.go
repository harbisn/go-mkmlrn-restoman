package models

import (
	"time"
)

type Room struct {
	ID           uint64    `json:"id" pg:"id, pk"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	Capacity     int8      `json:"capacity"`
	Description  string    `json:"description"`
	Type         string    `json:"type"`
	PricePerHour int32     `json:"pricePerHour"`
	CreatedBy    string    `json:"createdBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedBy    string    `json:"UpdatedBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

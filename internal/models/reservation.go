package models

import "time"

type Reservation struct {
	ID           uint64    `json:"id" pg:"id, pk"`
	Name         string    `json:"name"`
	CustomerName string    `json:"customerName"`
	PhoneNumber  string    `json:"phoneNumber"`
	Email        string    `json:"email"`
	Status       string    `json:"status"`
	CreatedBy    string    `json:"createdBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedBy    string    `json:"UpdatedBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

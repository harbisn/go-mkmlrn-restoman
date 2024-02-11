package models

import "time"

type PaginationResponse struct {
	Content      interface{}            `json:"content"`
	TotalElement int                    `json:"totalElement"`
	Page         int                    `json:"page"`
	Size         int                    `json:"size"`
	Order        string                 `json:"order"`
	Filter       map[string]interface{} `json:"filter"`
}

type RoomReservationDto struct {
	Name         string    `json:"name"`
	CustomerName string    `json:"customerName"`
	PhoneNumber  string    `json:"phoneNumber"`
	Email        string    `json:"email"`
	RoomID       uint64    `json:"roomId" pg:"fk"`
	RoomName     string    `json:"roomName"`
	Attendee     int8      `json:"attendee"`
	StartAt      time.Time `json:"StartAt"`
	EndAt        time.Time `json:"endAt"`
	Hours        int       `json:"hours"`
	Price        int32     `json:"price"`
}

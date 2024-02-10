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
	PricePerHour int32     `json:"pricePerHour"`
	CreatedBy    string    `json:"createdBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedBy    string    `json:"UpdatedBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (r *Room) CreateRoom() *Room {
	return Create(r).(*Room)
}

func GetAllRoom(offset, size int, order string, filters map[string]interface{}) ([]Room, error) {
	var rooms []Room
	if err := GetAll(&rooms, offset, size, order, filters); err != nil {
		return nil, err
	}
	return rooms, nil
}

func GetRoomById(ID uint64) *Room {
	var getRoom Room
	GetById(&getRoom, ID)
	return &getRoom
}

func (r *Room) UpdateRoom() *Room {
	Update(r)
	return r
}

func DeleteRoom(ID uint64) *Room {
	var deletedRoom Room
	Delete(&deletedRoom, ID)
	return &deletedRoom
}

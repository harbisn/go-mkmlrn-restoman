package models

import (
	"time"
)

type ReservationRoom struct {
	ReservationID   uint64    `json:"reservationId"`
	RoomID          uint64    `json:"roomId"`
	ReservationCode string    `json:"reservationCode"`
	StartAt         time.Time `json:"StartAt"`
	EndAt           time.Time `json:"endAt"`
	Attendee        int8      `json:"attendee"`
	Hours           int       `json:"hours"`
	Price           int32     `json:"price"`
	Status          string    `json:"status"`
	CreatedBy       string    `json:"createdBy"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedBy       string    `json:"UpdatedBy"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (reservationRoom *ReservationRoom) CreateReservationRoom() *ReservationRoom {
	return Create(reservationRoom).(*ReservationRoom)
}

func GetAllReservationRooms(offset, size int, order string, filters map[string]interface{}) ([]ReservationRoom, error) {
	var reservationRooms []ReservationRoom
	if err := GetAll(&reservationRooms, offset, size, order, filters); err != nil {
		return nil, err
	}
	return reservationRooms, nil
}

func (reservationRoom *ReservationRoom) UpdateReservationRoom() *ReservationRoom {
	Update(reservationRoom)
	return reservationRoom
}

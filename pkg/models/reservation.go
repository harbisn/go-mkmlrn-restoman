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

func (reservation *Reservation) CreateReservation() *Reservation {
	return Create(reservation).(*Reservation)
}

func GetAllReservations(offset, size int, order string, filters map[string]interface{}) ([]Reservation, error) {
	var reservations []Reservation
	if err := GetAll(&reservations, offset, size, order, filters); err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationByID(ID uint64) *Reservation {
	var getReservation Reservation
	GetById(&getReservation, ID)
	return &getReservation
}

func (reservation *Reservation) UpdateReservation() *Reservation {
	Update(reservation)
	return reservation
}

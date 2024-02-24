package reservation

import (
	"fmt"
	restomantime "github.com/harbisn/go-mkmlrn-restoman/internal/helper/time"
	"strconv"
	"time"
)

type Reservation struct {
	ID              uint64    `json:"id" pg:"id, pk"`
	Name            string    `json:"name"`
	CustomerName    string    `json:"customerName"`
	PhoneNumber     string    `json:"phoneNumber"`
	Email           string    `json:"email"`
	RoomID          uint64    `json:"roomId"`
	Attendee        int8      `json:"attendee"`
	StartAt         time.Time `json:"StartAt"`
	EndAt           time.Time `json:"endAt"`
	Hours           int       `json:"hours"`
	Price           int32     `json:"price"`
	ReservationCode string    `json:"reservationCode"`
	CreatedBy       string    `json:"createdBy"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedBy       string    `json:"UpdatedBy"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type CreateRequestDto struct {
	Name         string `json:"name"`
	CustomerName string `json:"customerName"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	RoomID       uint64 `json:"roomId"`
	Attendee     int8   `json:"attendee"`
	StartAt      string `json:"StartAt"`
	EndAt        string `json:"endAt"`
}

func CalculateHours(startAt, endAt time.Time) int {
	return int(endAt.Sub(startAt).Hours())
}

func CalculatePrice(hours int, pricePerHours int32) int32 {
	return int32(hours) * pricePerHours
}

func GenerateReservationCode(roomId uint64) string {
	strRoomId := strconv.Itoa(int(roomId))
	currentTime := restomantime.GetCurrentTime()
	return "R" + strRoomId + strconv.Itoa(currentTime.Nanosecond())
}

func ValidateBookingTime(startAt, endAt time.Time) error {
	if startAt.Hour() < 16 || endAt.Hour() > 22 {
		return fmt.Errorf("reservation outside operational hours")
	}
	return nil
}

func ValidateAttendee(attendee, capacity int8) error {
	if attendee > capacity {
		return fmt.Errorf("room overcapacity")
	}
	return nil
}

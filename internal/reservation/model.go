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
	StartAt         time.Time `pg:"default:now()" json:"startAt"`
	EndAt           time.Time `pg:"default:now()" json:"endAt"`
	Hours           int       `json:"hours"`
	Price           int32     `json:"price"`
	ReservationCode string    `json:"reservationCode"`
	CreatedBy       string    `json:"createdBy"`
	CreatedAt       time.Time `pg:"default:now()" json:"createdAt"`
	UpdatedBy       string    `json:"UpdatedBy"`
	UpdatedAt       time.Time `pg:"default:now()" json:"updatedAt"`
}

type CreateReservationRequestDto struct {
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
	if startAt.Hour() < 16 || endAt.Hour() > 23 {
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

// MapFromRequestWithMetaData set data from request and metadata
func MapFromRequestWithMetaData(requestDto *CreateReservationRequestDto, userId string) *Reservation {
	currentTime := restomantime.GetCurrentTime()
	return &Reservation{
		Name:            requestDto.Name,
		CustomerName:    requestDto.CustomerName,
		PhoneNumber:     requestDto.PhoneNumber,
		Email:           requestDto.Email,
		RoomID:          requestDto.RoomID,
		Attendee:        requestDto.Attendee,
		ReservationCode: GenerateReservationCode(requestDto.RoomID),
		CreatedBy:       userId,
		CreatedAt:       currentTime,
		UpdatedBy:       userId,
		UpdatedAt:       currentTime,
	}
}

// MapToWithLocalTime return new instance for immutability and convert time to local indonesian
func MapToWithLocalTime(reservationFromDB Reservation) Reservation {
	location := restomantime.GetLocation()
	return Reservation{
		ID:              reservationFromDB.ID,
		Name:            reservationFromDB.Name,
		CustomerName:    reservationFromDB.CustomerName,
		PhoneNumber:     reservationFromDB.PhoneNumber,
		Email:           reservationFromDB.Email,
		RoomID:          reservationFromDB.RoomID,
		Attendee:        reservationFromDB.Attendee,
		StartAt:         reservationFromDB.StartAt.In(location),
		EndAt:           reservationFromDB.EndAt.In(location),
		Hours:           reservationFromDB.Hours,
		Price:           reservationFromDB.Price,
		ReservationCode: reservationFromDB.ReservationCode,
		CreatedBy:       reservationFromDB.CreatedBy,
		CreatedAt:       reservationFromDB.CreatedAt.In(location),
		UpdatedBy:       reservationFromDB.UpdatedBy,
		UpdatedAt:       reservationFromDB.UpdatedAt.In(location),
	}
}

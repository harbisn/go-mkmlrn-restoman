package handlers

import (
	"github.com/harbisn/go-mkmlrn-restoman/pkg/models"
	"net/http"
	"strconv"
)

func MakeRoomReservation(w http.ResponseWriter, r *http.Request) {
	RoomReservationDto := &models.RoomReservationDto{}
	ParseJSONRequestBody(r, RoomReservationDto)
	userId := r.Header.Get("x-user-id")

	CreateReservation := &models.Reservation{}
	CreateReservation.Name = RoomReservationDto.Name
	CreateReservation.CustomerName = RoomReservationDto.CustomerName
	CreateReservation.PhoneNumber = RoomReservationDto.PhoneNumber
	CreateReservation.Email = RoomReservationDto.Email
	CreateReservation.Status = "WAITING_PAYMENT"
	CreateReservation.CreatedBy = userId
	CreateReservation.UpdatedBy = userId
	createdReservation := CreateReservation.CreateReservation()

	reservedRoom := models.GetRoomByID(RoomReservationDto.RoomID)
	if reservedRoom.Status != "AVAILABLE" {
		// TODO: send response --> "room already booked"
		return
	}
	if RoomReservationDto.Attendee > reservedRoom.Capacity {
		// TODO: send response --> "reserve room with larger capacity"
		return
	}

	hours := RoomReservationDto.EndAt.Sub(RoomReservationDto.StartAt).Hours()
	price := reservedRoom.PricePerHour * int32(hours)

	CreateReservationRoom := &models.ReservationRoom{}
	CreateReservationRoom.ReservationID = createdReservation.ID
	CreateReservationRoom.RoomID = reservedRoom.ID
	CreateReservationRoom.ReservationCode = GetReservationCode(reservedRoom.ID, createdReservation.ID)
	CreateReservationRoom.StartAt = RoomReservationDto.StartAt
	CreateReservationRoom.EndAt = RoomReservationDto.EndAt
	CreateReservationRoom.Attendee = RoomReservationDto.Attendee
	CreateReservationRoom.Hours = int(hours)
	CreateReservationRoom.Price = price
	CreateReservationRoom.Status = "WAITING_PAYMENT"
	createdReservationRoom := CreateReservationRoom.CreateReservationRoom()

	RoomReservationDto.RoomName = reservedRoom.Name
	RoomReservationDto.Hours = createdReservationRoom.Hours
	RoomReservationDto.Price = createdReservationRoom.Price
	SendJSONResponse(w, http.StatusCreated, RoomReservationDto)
}

func GetReservationCode(roomId, reservationId uint64) string {
	strRoomId := strconv.Itoa(int(roomId))
	strReservationId := strconv.Itoa(int(reservationId))
	currentTime := models.GetCurrentTime()
	ss := currentTime.Second()
	mmm := currentTime.Nanosecond() / 1000000
	return "R" + strRoomId + strReservationId + strconv.Itoa(ss) + strconv.Itoa(mmm)
}

func AddMoreRooms(w http.ResponseWriter, r *http.Request) {
	CreateReservationRoom := &models.ReservationRoom{}
	ParseJSONRequestBody(r, CreateReservationRoom)

	reservedRoom := models.GetRoomByID(CreateReservationRoom.RoomID)
	if reservedRoom.Status != "AVAILABLE" {
		// TODO: send response --> "room already booked"
		return
	}
	if CreateReservationRoom.Attendee > reservedRoom.Capacity {
		// TODO: send response --> "reserve room with larger capacity"
		return
	}

	hours := CreateReservationRoom.EndAt.Sub(CreateReservationRoom.StartAt).Hours()
	price := reservedRoom.PricePerHour * int32(hours)

	userId := r.Header.Get("x-user-id")
	CreateReservationRoom.CreatedBy = userId
	CreateReservationRoom.UpdatedBy = userId
	CreateReservationRoom.Hours = int(hours)
	CreateReservationRoom.Price = price
	createdReservationRoom := CreateReservationRoom.CreateReservationRoom()

	SendJSONResponse(w, http.StatusCreated, createdReservationRoom)
}

func GetAllReservationRooms(w http.ResponseWriter, r *http.Request) {
	params := []string{"reservationId", "roomId", "reservationCode", "status"}
	offset, size, order, filters := GetFilterAndPagination(r, params)
	reservationRooms, _ := models.GetAllReservationRooms(offset, size, order, filters)
	pageReservationRooms := Paginate(reservationRooms, len(reservationRooms), size, offset, order, filters)
	SendJSONResponse(w, http.StatusOK, pageReservationRooms)
}

// TODO: cancel reservation_room --> validate if other room canceled or not, if true then cancel the reservation
// TODO: update reservation_room, validate if already paid or DP can't update hours, validate attendee so that doesn't over capacity

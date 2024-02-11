package handlers

import (
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
	"github.com/harbisn/go-mkmlrn-restoman/internal/repository"
	"net/http"
	"strconv"
	"time"
)

type ReservationRoomHandler struct {
	reservationRepository     *repository.ReservationRepository
	reservationRoomRepository *repository.ReservationRoomRepository
	roomRepository            *repository.RoomRepository
}

func NewReservationRoomHandler(re *repository.ReservationRepository,
	rR *repository.ReservationRoomRepository, r *repository.RoomRepository) *ReservationRoomHandler {
	return &ReservationRoomHandler{
		reservationRepository:     re,
		reservationRoomRepository: rR,
		roomRepository:            r,
	}
}

func (h *ReservationRoomHandler) MakeRoomReservation(w http.ResponseWriter, r *http.Request) {
	roomReservationDto := &models.RoomReservationDto{}
	ParseJSONRequestBody(r, roomReservationDto)
	userId := r.Header.Get("x-user-id")

	createReservation := &models.Reservation{}
	createReservation.Name = roomReservationDto.Name
	createReservation.CustomerName = roomReservationDto.CustomerName
	createReservation.PhoneNumber = roomReservationDto.PhoneNumber
	createReservation.Email = roomReservationDto.Email
	createReservation.Status = "WAITING_PAYMENT"
	createReservation.CreatedBy = userId
	createReservation.UpdatedBy = userId
	createdReservation := h.reservationRepository.Create(createReservation)

	reservedRoom := h.roomRepository.GetById(roomReservationDto.RoomID)
	if reservedRoom.Status != "AVAILABLE" {
		// TODO: send response --> "room already booked"
		return
	}
	if roomReservationDto.Attendee > reservedRoom.Capacity {
		// TODO: send response --> "reserve room with larger capacity"
		return
	}

	hours := roomReservationDto.EndAt.Sub(roomReservationDto.StartAt).Hours()
	price := reservedRoom.PricePerHour * int32(hours)

	createReservationRoom := &models.ReservationRoom{}
	createReservationRoom.ReservationID = createdReservation.ID
	createReservationRoom.RoomID = reservedRoom.ID
	createReservationRoom.ReservationCode = GetReservationCode(reservedRoom.ID, createdReservation.ID)
	createReservationRoom.StartAt = roomReservationDto.StartAt
	createReservationRoom.EndAt = roomReservationDto.EndAt
	createReservationRoom.Attendee = roomReservationDto.Attendee
	createReservationRoom.Hours = int(hours)
	createReservationRoom.Price = price
	createReservationRoom.Status = "WAITING_PAYMENT"
	createdReservationRoom := h.reservationRoomRepository.Create(createReservationRoom)

	roomReservationDto.RoomName = reservedRoom.Name
	roomReservationDto.Hours = createdReservationRoom.Hours
	roomReservationDto.Price = createdReservationRoom.Price
	SendJSONResponse(w, http.StatusCreated, roomReservationDto)
}

func GetReservationCode(roomId, reservationId uint64) string {
	strRoomId := strconv.Itoa(int(roomId))
	strReservationId := strconv.Itoa(int(reservationId))
	location, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now().UTC().In(location)
	ss := currentTime.Second()
	mmm := currentTime.Nanosecond() / 1000000
	return "R" + strRoomId + strReservationId + strconv.Itoa(ss) + strconv.Itoa(mmm)
}

func (h *ReservationRoomHandler) AddMoreRooms(w http.ResponseWriter, r *http.Request) {
	createReservationRoom := &models.ReservationRoom{}
	ParseJSONRequestBody(r, createReservationRoom)

	reservedRoom := h.roomRepository.GetById(createReservationRoom.RoomID)
	if reservedRoom.Status != "AVAILABLE" {
		// TODO: send response --> "room already booked"
		return
	}
	if createReservationRoom.Attendee > reservedRoom.Capacity {
		// TODO: send response --> "reserve room with larger capacity"
		return
	}

	hours := createReservationRoom.EndAt.Sub(createReservationRoom.StartAt).Hours()
	price := reservedRoom.PricePerHour * int32(hours)

	userId := r.Header.Get("x-user-id")
	createReservationRoom.CreatedBy = userId
	createReservationRoom.UpdatedBy = userId
	createReservationRoom.Hours = int(hours)
	createReservationRoom.Price = price
	createdReservationRoom := h.reservationRoomRepository.Create(createReservationRoom)

	SendJSONResponse(w, http.StatusCreated, createdReservationRoom)
}

func (h *ReservationRoomHandler) GetAllReservationRooms(w http.ResponseWriter, r *http.Request) {
	params := []string{"reservationId", "roomId", "reservationCode", "status"}
	offset, size, order, filters := GetFilterAndPagination(r, params)
	var p = models.PageableDto{}
	p.Offset = offset
	p.Size = size
	p.Order = order
	p.Filter = filters
	reservationRooms, _ := h.reservationRoomRepository.GetAll(p)
	pageReservationRooms := Paginate(reservationRooms, len(reservationRooms), size, offset, order, filters)
	SendJSONResponse(w, http.StatusOK, pageReservationRooms)
}

// TODO: cancel reservation_room --> validate if other room canceled or not, if true then cancel the reservation
// TODO: update reservation_room, validate if already paid or DP can't update hours, validate attendee so that doesn't over capacity

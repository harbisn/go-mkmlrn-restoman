package handlers

import (
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
	"github.com/harbisn/go-mkmlrn-restoman/internal/repository"
	"net/http"
)

type RoomHandler struct {
	roomRepository *repository.RoomRepository
}

func NewRoomHandler(r *repository.RoomRepository) *RoomHandler {
	return &RoomHandler{roomRepository: r}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	room := &models.Room{}
	ParseJSONRequestBody(r, room)
	userId := r.Header.Get("x-user-id")
	room.CreatedBy = userId
	room.UpdatedBy = userId
	room = h.roomRepository.Create(room)
	SendJSONResponse(w, http.StatusCreated, room)
}

func (h *RoomHandler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	params := []string{"status", "capacity", "highestPricePerHour", "lowestPricePerHour"}
	offset, size, order, filters := GetFilterAndPagination(r, params)
	var p = models.PageableDto{}
	p.Offset = offset
	p.Size = size
	p.Order = order
	p.Filter = filters
	rooms, _ := h.roomRepository.GetAll(p)
	pageRooms := Paginate(rooms, len(rooms), size, offset, order, filters)
	SendJSONResponse(w, http.StatusOK, pageRooms)
}

func (h *RoomHandler) GetRoomByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	room := h.roomRepository.GetById(id)
	SendJSONResponse(w, http.StatusOK, room)
}

func (h *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	room := h.roomRepository.GetById(id)
	updateRoom := &models.Room{}
	ParseJSONRequestBody(r, updateRoom)
	if updateRoom.Status != "" {
		room.Status = updateRoom.Status
	}
	if updateRoom.Type != "" {
		room.Type = updateRoom.Type
	}
	if updateRoom.Capacity != room.Capacity {
		room.Capacity = updateRoom.Capacity
	}
	if updateRoom.PricePerHour != room.PricePerHour {
		room.PricePerHour = updateRoom.PricePerHour
	}
	room.UpdatedBy = r.Header.Get("x-user-id")
	m := h.roomRepository.Update(updateRoom)
	SendJSONResponse(w, http.StatusOK, m)
}

func (h *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	room := h.roomRepository.Delete(id)
	SendJSONResponse(w, http.StatusOK, room)
}

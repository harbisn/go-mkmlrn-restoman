package room

import (
	"encoding/json"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/handler"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
	"net/http"
)

type Handler struct {
	roomService *Service
}

func NewRoomHandler(s *Service) *Handler {
	return &Handler{roomService: s}
}

const BadRequestMessage = "Request can't be parsed."

func (h *Handler) InsertRoomHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var room Room
	err := decoder.Decode(&room)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusBadRequest, BadRequestMessage)
		return
	}
	userId := r.Header.Get("x-user-id")
	err = h.roomService.InsertRoom(&room, userId)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusInternalServerError, "Failed to insert new room.")
		return
	}
	handler.WriteSuccessResponse(w, http.StatusCreated, room, nil)
	return
}

func (h *Handler) SelectRoomHandler(w http.ResponseWriter, r *http.Request) {
	params := []string{"status", "capacity", "highestPricePerHour", "lowestPricePerHour"}
	pageableDto := pagination.GetFilterAndPagination(r, params)
	pageRoom, err := h.roomService.SelectRoom(pageableDto)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusInternalServerError, "Failed to select room.")
		return
	}
	handler.WriteSuccessResponse(w, http.StatusOK, pageRoom, nil)
}

func (h *Handler) UpdateRoomHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var room Room
	err := decoder.Decode(&room)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusBadRequest, BadRequestMessage)
		return
	}
	id, err := handler.GetIdFromPath(r)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusBadRequest, BadRequestMessage)
		return
	}
	userId := r.Header.Get("x-user-id")
	updatedRoom, err := h.roomService.UpdateRoom(id, &room, userId)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusInternalServerError, "Failed to update room.")
		return
	}
	handler.WriteSuccessResponse(w, http.StatusCreated, updatedRoom, nil)
}

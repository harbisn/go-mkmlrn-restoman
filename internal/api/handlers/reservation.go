package handlers

import (
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
	"github.com/harbisn/go-mkmlrn-restoman/internal/repository"
	"net/http"
)

type ReservationHandler struct {
	reservationRepository *repository.ReservationRepository
}

func NewReservationHandler(r *repository.ReservationRepository) *ReservationHandler {
	return &ReservationHandler{reservationRepository: r}
}

func (h *ReservationHandler) GetAllReservations(w http.ResponseWriter, r *http.Request) {
	params := []string{"customerName", "phoneNumber", "email"}
	offset, size, order, filters := GetFilterAndPagination(r, params)
	var p = models.PageableDto{}
	p.Offset = offset
	p.Size = size
	p.Order = order
	p.Filter = filters
	reservations, _ := h.reservationRepository.GetAll(p)
	pageReservations := Paginate(reservations, len(reservations), size, offset, order, filters)
	SendJSONResponse(w, http.StatusOK, pageReservations)
}

func (h *ReservationHandler) GetReservationByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	reservationDetails := h.reservationRepository.GetById(id)
	SendJSONResponse(w, http.StatusOK, reservationDetails)
}

func (h *ReservationHandler) UpdateReservation(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	reservationDetails := h.reservationRepository.GetById(id)
	updateReservation := &models.Reservation{}
	ParseJSONRequestBody(r, updateReservation)
	if updateReservation.Name != "" {
		reservationDetails.Name = updateReservation.Name
	}
	if updateReservation.CustomerName != "" {
		reservationDetails.CustomerName = updateReservation.CustomerName
	}
	if updateReservation.PhoneNumber != "" {
		reservationDetails.PhoneNumber = updateReservation.PhoneNumber
	}
	if updateReservation.Email != "" {
		reservationDetails.Email = updateReservation.Email
	}
	reservationDetails.UpdatedBy = r.Header.Get("x-user-id")
	re := h.reservationRepository.Update(reservationDetails)
	SendJSONResponse(w, http.StatusOK, re)
}

// TODO: cancel reservation --> automatically cancel all reservation_room

package handlers

import (
	"github.com/harbisn/go-mkmlrn-restoman/pkg/models"
	"net/http"
)

func GetAllReservations(w http.ResponseWriter, r *http.Request) {
	params := []string{"customerName", "phoneNumber", "email"}
	offset, size, order, filters := GetFilterAndPagination(r, params)
	reservations, _ := models.GetAllReservations(offset, size, order, filters)
	pageReservations := Paginate(reservations, len(reservations), size, offset, order, filters)
	SendJSONResponse(w, http.StatusOK, pageReservations)
}

func GetReservationByID(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	reservationDetails := models.GetReservationByID(id)
	SendJSONResponse(w, http.StatusOK, reservationDetails)
}

func UpdateReservation(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIDFromRequestToUint64(r, "id")
	ValidateInternalError(w, err)
	reservationDetails := models.GetReservationByID(id)
	UpdateReservation := &models.Reservation{}
	ParseJSONRequestBody(r, UpdateReservation)
	if UpdateReservation.Name != "" {
		reservationDetails.Name = UpdateReservation.Name
	}
	if UpdateReservation.CustomerName != "" {
		reservationDetails.CustomerName = UpdateReservation.CustomerName
	}
	if UpdateReservation.PhoneNumber != "" {
		reservationDetails.PhoneNumber = UpdateReservation.PhoneNumber
	}
	if UpdateReservation.Email != "" {
		reservationDetails.Email = UpdateReservation.Email
	}
	reservationDetails.UpdatedBy = r.Header.Get("x-user-id")
	m := reservationDetails.UpdateReservation()
	SendJSONResponse(w, http.StatusOK, m)
}

// TODO: cancel reservation --> automatically cancel all reservation_room

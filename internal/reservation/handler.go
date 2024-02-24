package reservation

import (
	"encoding/json"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/handler"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

type Handler struct {
	reservationService *Service
}

func NewReservationHandler(s *Service) *Handler {
	return &Handler{reservationService: s}
}

const BadRequestMessage = "Request can't be parsed."

func (h *Handler) CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	decoder := json.NewDecoder(r.Body)
	var requestDto InsertRequestDto
	err := decoder.Decode(&requestDto)
	if err != nil {
		log.Error().Err(err).Msgf("error while parsing request %s", err.Error())
		handler.WriteFailResponse(w, http.StatusBadRequest, BadRequestMessage)
		return
	}
	userId := r.Header.Get("x-user-id")
	var reservation Reservation
	err = h.reservationService.CreateReservation(&requestDto, &reservation, userId)
	if err != nil {
		handler.WriteFailResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.WriteSuccessResponse(w, http.StatusCreated, reservation, nil)
	return
}

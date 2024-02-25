package reservation

import (
	"fmt"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
	restomantime "github.com/harbisn/go-mkmlrn-restoman/internal/helper/time"
	"github.com/harbisn/go-mkmlrn-restoman/internal/room"
	"github.com/rs/zerolog"
)

type Service struct {
	reservationRepository *Repository
	roomRepository        *room.Repository
	logger                zerolog.Logger
}

func NewReservationService(r *Repository, roomRepository *room.Repository, logger zerolog.Logger) *Service {
	return &Service{reservationRepository: r, roomRepository: roomRepository, logger: logger}
}

func (s *Service) CreateReservation(requestDto *CreateReservationRequestDto, userId string) (*Reservation, error) {
	logger := s.logger.With().Timestamp().Logger()

	reservation := MapFromRequestWithMetaData(requestDto, userId)

	if err := s.setReservationTime(requestDto, reservation); err != nil {
		return nil, err
	}

	selectedRoom, err := s.validateReservation(reservation)
	if err != nil {
		return nil, err
	}

	reservation.Hours = CalculateHours(reservation.StartAt, reservation.EndAt)
	reservation.Price = CalculatePrice(reservation.Hours, selectedRoom.PricePerHour)

	if err := s.reservationRepository.Insert(reservation); err != nil {
		logger.Error().Err(err).Msg("failed to insert new reservation")
		return nil, fmt.Errorf("failed to insert new reservation: %w", err)
	}

	return reservation, nil
}

func (s *Service) setReservationTime(requestDto *CreateReservationRequestDto, reservation *Reservation) error {
	startAt, err := restomantime.StrToLocalTime(requestDto.StartAt)
	if err != nil {
		s.logger.Error().Msg(err.Error())
		return err
	}
	reservation.StartAt = startAt

	endAt, err := restomantime.StrToLocalTime(requestDto.EndAt)
	if err != nil {
		s.logger.Error().Msg(err.Error())
		return err
	}
	reservation.EndAt = endAt

	return nil
}

func (s *Service) validateReservation(reservation *Reservation) (*room.Room, error) {
	roomId := reservation.RoomID
	startAt := reservation.StartAt
	endAt := reservation.EndAt

	if err := ValidateBookingTime(startAt, endAt); err != nil {
		s.logger.Error().Msg(err.Error())
		return &room.Room{}, err
	}

	selectedRoom, err := s.roomRepository.SelectById(roomId)
	if err != nil {
		s.logger.Error().Err(err).Msg("error selecting room")
		return &room.Room{}, fmt.Errorf("error selecting room: %w", err)
	}

	if err := ValidateAttendee(reservation.Attendee, selectedRoom.Capacity); err != nil {
		s.logger.Error().Msg(err.Error())
		return &room.Room{}, err
	}

	existingReservations, err := s.reservationRepository.FindReservationsWithinTimeRange(roomId, startAt, endAt)
	if err != nil {
		s.logger.Error().Err(err).Msg("error checking room availability")
		return &room.Room{}, fmt.Errorf("error checking room availability: %w", err)
	}
	if existingReservations != nil {
		s.logger.Error().Msg("room not available / already booked")
		return &room.Room{}, fmt.Errorf("room not available / already booked")
	}

	return selectedRoom, nil
}

func (s *Service) GetReservations(pageable pagination.PageableDto) (pagination.PageableDto, error) {
	reservations, count, err := s.reservationRepository.Select(pageable)
	if err != nil {
		s.logger.Err(err).Msg("Failed to select menu")
		return pagination.PageableDto{}, err
	}
	pageable.TotalElement = count
	pageable.NumberOfElement = len(reservations)
	pageReservation := pagination.Paginate(reservations, pageable)
	return pageReservation, nil
}

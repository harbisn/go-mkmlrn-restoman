package reservation

import (
	"fmt"
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

func (s *Service) CreateReservation(requestDto *CreateRequestDto, reservation *Reservation, userId string) error {
	logger := s.logger.With().Timestamp().Logger()
	currentTime := restomantime.GetCurrentTime()

	reservation.Name = requestDto.Name
	reservation.CustomerName = requestDto.CustomerName
	reservation.PhoneNumber = requestDto.PhoneNumber
	reservation.RoomID = requestDto.RoomID
	reservation.Attendee = requestDto.Attendee
	reservation.CreatedBy = userId
	reservation.CreatedAt = currentTime
	reservation.UpdatedBy = userId
	reservation.UpdatedAt = currentTime
	reservation.ReservationCode = GenerateReservationCode(reservation.RoomID)

	if err := s.setReservationTime(requestDto, reservation); err != nil {
		return err
	}

	if err := s.validateReservation(reservation); err != nil {
		return err
	}

	if err := s.reservationRepository.Insert(reservation); err != nil {
		logger.Error().Err(err).Msg("failed to insert new reservation")
		return fmt.Errorf("failed to insert new reservation: %w", err)
	}

	return nil
}

func (s *Service) setReservationTime(requestDto *CreateRequestDto, reservation *Reservation) error {
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

func (s *Service) validateReservation(reservation *Reservation) error {
	roomId := reservation.RoomID
	startAt := reservation.StartAt
	endAt := reservation.EndAt

	if err := ValidateBookingTime(startAt, endAt); err != nil {
		s.logger.Error().Msg(err.Error())
		return err
	}

	selectedRoom, err := s.roomRepository.SelectById(roomId)
	if err != nil {
		s.logger.Error().Err(err).Msg("error selecting room")
		return fmt.Errorf("error selecting room: %w", err)
	}

	reservation.Hours = CalculateHours(startAt, endAt)
	reservation.Price = CalculatePrice(reservation.Hours, selectedRoom.PricePerHour)

	if err := ValidateAttendee(reservation.Attendee, selectedRoom.Capacity); err != nil {
		s.logger.Error().Msg(err.Error())
		return err
	}

	existingReservations, err := s.reservationRepository.FindReservationsWithinTimeRange(roomId, startAt, endAt)
	if err != nil {
		s.logger.Error().Err(err).Msg("error checking room availability")
		return fmt.Errorf("error checking room availability: %w", err)
	}
	if existingReservations != nil {
		s.logger.Error().Msg("room not available / already booked")
		return fmt.Errorf("room not available / already booked")
	}

	return nil
}

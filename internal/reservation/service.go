package reservation

import (
	"fmt"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
	restomantime "github.com/harbisn/go-mkmlrn-restoman/internal/helper/time"
	"github.com/harbisn/go-mkmlrn-restoman/internal/room"
	"github.com/rs/zerolog"
	"time"
)

type Service struct {
	reservationRepository *Repository
	roomRepository        *room.Repository
	logger                zerolog.Logger
}

func NewReservationService(r *Repository, roomRepository *room.Repository, logger zerolog.Logger) *Service {
	return &Service{reservationRepository: r, roomRepository: roomRepository, logger: logger}
}

func (s *Service) CreateReservation(requestDto *RequestDto, userId string) (*Reservation, error) {
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

func (s *Service) setReservationTime(requestDto *RequestDto, reservation *Reservation) error {
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

// GetAvailableTimes returns a list of available times for a room starting from the specified time.
func (s *Service) GetAvailableTimes(roomId uint64, strStartAt string) ([]AvailableTimeDto, error) {
	startAt, err := restomantime.StrToLocalTime(strStartAt)
	if err != nil {
		s.logger.Error().Msg(err.Error())
		return nil, err
	}

	existingReservation, err := s.reservationRepository.FindByDate(roomId, startAt)
	if err != nil {
		return nil, err
	}

	var availableTimes []AvailableTimeDto
	n := len(existingReservation)
	soh, eoh := GetOperationalHours(startAt)
	minGap := 90 * time.Minute

	// if n == 1, compare startAt with 16 and endAt with 23, if >= 90 minutes then display as available times
	if n == 1 {
		availableTime := s.getAvailableTimesForSingleReservation(existingReservation, soh, eoh, minGap)
		availableTimes = append(availableTimes, availableTime)
	}

	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			availableTimes = s.getAvailableTimesForInterval(existingReservation, i, n, soh, eoh, minGap)
		}
	}

	if n%2 != 0 {
		for i := 0; i < n-1; i++ {
			if i == 0 {
				availableTime := s.getAvailableTimesForSingleReservation(existingReservation, soh, eoh, minGap)
				availableTimes = append(availableTimes, availableTime)
			} else {
				availableTimes = s.getAvailableTimesForInterval(existingReservation, i, n, soh, eoh, minGap)
			}
		}
	}

	return availableTimes, nil
}

// getAvailableTimesForSingleReservation calculates available times when there's only one existing reservation.
func (s *Service) getAvailableTimesForSingleReservation(existingReservation []Reservation,
	soh, eoh time.Time, minGap time.Duration) AvailableTimeDto {
	ers := existingReservation[0].StartAt
	ere := existingReservation[0].EndAt

	var availableTime AvailableTimeDto
	roomId := existingReservation[0].RoomID
	m15 := 15 * time.Minute

	if ers.Sub(soh) >= minGap {
		availableTime = SetAvailableTime(roomId, soh, ers.Add(-m15))
	}
	if eoh.Sub(ere) >= minGap {
		availableTime = SetAvailableTime(roomId, ere.Add(m15), eoh)
	}
	return availableTime
}

// getAvailableTimesForInterval calculates available times between two existing reservations.
func (s *Service) getAvailableTimesForInterval(existingReservation []Reservation, i, n int,
	soh, eoh time.Time, minGap time.Duration) []AvailableTimeDto {
	var availableTimes []AvailableTimeDto
	m15 := 15 * time.Minute
	roomId := existingReservation[i].RoomID

	// if first startAt > 16
	fErs := existingReservation[i].StartAt
	if i == 0 && fErs.Sub(soh) >= minGap {
		availableTimes = append(availableTimes, SetAvailableTime(roomId, soh, fErs.Add(-m15)))
	}

	// if last endAt < 23
	lEre := existingReservation[i+1].EndAt
	if i+1 == n-1 && eoh.Sub(lEre) >= minGap {
		availableTimes = append(availableTimes, SetAvailableTime(roomId, lEre.Add(m15), eoh))
	}

	t0 := existingReservation[i].EndAt
	t1 := existingReservation[i+1].StartAt
	gap := t1.Sub(t0)
	if gap >= minGap {
		availableTimes = append(availableTimes, SetAvailableTime(roomId, t0.Add(m15), t1.Add(-m15)))
	}

	return availableTimes
}

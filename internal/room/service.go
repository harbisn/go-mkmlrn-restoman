package room

import (
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/time"
	"github.com/rs/zerolog"
	"os"
)

type Service struct {
	roomRepository *Repository
}

func NewRoomService(r *Repository) *Service {
	return &Service{roomRepository: r}
}

func (s *Service) InsertRoom(room *Room, userId string) error {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	var currentTime = time.GetCurrentTime()
	room.CreatedBy = userId
	room.CreatedAt = currentTime
	room.UpdatedBy = userId
	room.UpdatedAt = currentTime
	err := s.roomRepository.Insert(room)
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert new room")
	}
	return err
}

func (s *Service) SelectRoom(pageable pagination.PageableDto) (pagination.PageableDto, error) {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	rooms, count, err := s.roomRepository.Select(pageable)
	if err != nil {
		log.Error().Err(err).Msg("Failed to select room")
		return pagination.PageableDto{}, err
	}
	pageable.TotalElement = count
	pageable.NumberOfElement = len(rooms)
	pageRoom := pagination.Paginate(rooms, pageable)
	return pageRoom, nil
}

func (s *Service) SelectRoomById(id uint64) (*Room, error) {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	room, err := s.roomRepository.SelectById(id)
	if err != nil {
		log.Error().Err(err).Msgf("Room with id %d is not found", id)
		return nil, err
	}
	return room, nil
}

func (s *Service) UpdateRoom(id uint64, updatedRoom *Room, userId string) (*Room, error) {
	currentRoom, err := s.SelectRoomById(id)
	if err != nil {
		return nil, err
	}
	if updatedRoom.Name != "" {
		currentRoom.Name = updatedRoom.Name
	}
	if updatedRoom.Status != "" {
		currentRoom.Status = updatedRoom.Status
	}
	if updatedRoom.Capacity != currentRoom.Capacity {
		currentRoom.Capacity = updatedRoom.Capacity
	}
	if updatedRoom.PricePerHour != currentRoom.PricePerHour {
		currentRoom.PricePerHour = updatedRoom.PricePerHour
	}
	if updatedRoom.Description != "" {
		currentRoom.Description = updatedRoom.Description
	}
	currentRoom.UpdatedBy = userId
	currentRoom.UpdatedAt = time.GetCurrentTime()
	if err := s.roomRepository.Update(currentRoom); err != nil {
		return nil, err
	}
	return currentRoom, nil
}

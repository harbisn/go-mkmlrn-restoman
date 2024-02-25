package menu

import (
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/time"
	"github.com/rs/zerolog"
	"os"
)

type Service struct {
	menuRepository *Repository
}

func NewMenuService(r *Repository) *Service {
	return &Service{menuRepository: r}
}

func (s *Service) CreateMenu(menu *Menu, userId string) error {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	var currentTime = time.GetCurrentTime()
	menu.CreatedBy = userId
	menu.CreatedAt = currentTime
	menu.UpdatedBy = userId
	menu.UpdatedAt = currentTime
	err := s.menuRepository.Insert(menu)
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert new menu")
	}
	return err
}

func (s *Service) GetMenus(pageable pagination.PageableDto) (pagination.PageableDto, error) {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	menus, count, err := s.menuRepository.Select(pageable)
	if err != nil {
		log.Error().Err(err).Msg("Failed to select menu")
		return pagination.PageableDto{}, err
	}
	pageable.TotalElement = count
	pageable.NumberOfElement = len(menus)
	pageMenu := pagination.Paginate(menus, pageable)
	return pageMenu, nil
}

func (s *Service) SelectMenuById(id uint64) (*Menu, error) {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	menu, err := s.menuRepository.SelectById(id)
	if err != nil {
		log.Error().Err(err).Msgf("Menu with id %d is not found", id)
		return nil, err
	}
	return menu, nil
}

func (s *Service) UpdateMenu(id uint64, updatedMenu *Menu, userId string) (*Menu, error) {
	currentMenu, err := s.SelectMenuById(id)
	if err != nil {
		return nil, err
	}
	if updatedMenu.Name != "" {
		currentMenu.Name = updatedMenu.Name
	}
	if updatedMenu.Status != "" {
		currentMenu.Status = updatedMenu.Status
	}
	if updatedMenu.Category != "" {
		currentMenu.Category = updatedMenu.Category
	}
	if updatedMenu.Price != currentMenu.Price {
		currentMenu.Price = updatedMenu.Price
	}
	if updatedMenu.Description != "" {
		currentMenu.Description = updatedMenu.Description
	}
	currentMenu.UpdatedBy = userId
	currentMenu.UpdatedAt = time.GetCurrentTime()
	if err := s.menuRepository.Update(currentMenu); err != nil {
		return nil, err
	}
	return currentMenu, nil
}

package models

import (
	"time"
)

type Menu struct {
	ID          uint64    `json:"id" pg:"id, pk"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Price       int32     `json:"price"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedBy   string    `json:"UpdatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (m *Menu) CreateMenu() *Menu {
	return Create(m).(*Menu)
}

func GetAllMenu(offset, size int, order string, filters map[string]interface{}) ([]Menu, error) {
	var menus []Menu
	if err := GetAll(&menus, offset, size, order, filters); err != nil {
		return nil, err
	}
	return menus, nil
}

func GetMenuById(ID uint64) *Menu {
	var getMenu Menu
	GetById(&getMenu, ID)
	return &getMenu
}

func (m *Menu) UpdateMenu() *Menu {
	Update(m)
	return m
}

func DeleteMenu(ID uint64) *Menu {
	var deletedMenu Menu
	Delete(&deletedMenu, ID)
	return &deletedMenu
}

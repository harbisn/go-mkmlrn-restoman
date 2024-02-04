package models

import (
	"github.com/harbisn/go-mkmlrn-restoman/pkg/database"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type Menu struct {
	ID          uint64    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Category    string    `json:"category"`
	Price       int32     `json:"price"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedBy   string    `json:"UpdatedBy"`
}

func init() {
	database.Connect()
	db = database.GetDB()
	// TODO: AutoMigrate will create new column if we changed names or type instead of update, it also won't drop deleted field from struct
	// TODO: Find away to validate table changes and do migration if necessary
	err := db.AutoMigrate(&Menu{})
	if err != nil {
		return
	}
}

func (m *Menu) CreateMenu() *Menu {
	db.Create(m)
	return m
}

func GetAllMenu() []Menu {
	var Menus []Menu
	db.Find(&Menus)
	return Menus
}

func GetMenuById(ID uint64) (*Menu, *gorm.DB) {
	var getMenu Menu
	db := db.Where("id = ?", ID).Find(&getMenu)
	return &getMenu, db
}

func (m *Menu) UpdateMenu(ID uint64) *Menu {
	m.ID = ID
	db.Save(m)
	return m
}

func DeleteMenu(ID uint64) Menu {
	var menu Menu
	db.Where("id = ?", ID).Delete(menu)
	return menu
}

package models

import (
	"github.com/harbisn/go-mkmlrn-restoman/pkg/database"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type Menu struct {
	ID          uint64    `json:"id" db:"id" gorm:"primary_key"`
	Name        string    `json:"name" db:"name"`
	Code        string    `json:"code" db:"code"`
	Status      string    `json:"status" db:"status"`
	Category    string    `json:"category" db:"category"`
	Price       int32     `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

func init() {
	database.Connect()
	db = database.GetDB()
	err := db.AutoMigrate(&Menu{})
	if err != nil {
		return
	}
}

func GetAllMenu() []Menu {
	var Menus []Menu
	db.Find(&Menus)
	return Menus
}

func GetMenuByCode(Code string) (*Menu, *gorm.DB) {
	var getMenu Menu
	db := db.Where("code = ?", Code).Find(&getMenu)
	return &getMenu, db
}

func (m *Menu) CreateMenu() *Menu {
	db.Create(m)
	return m
}

func (m *Menu) UpdateMenu(ID uint64) *Menu {
	m.ID = ID
	db.Save(m)
	return m
}

package models

import (
	"github.com/harbisn/go-mkmlrn-restoman/pkg/database"
	"gorm.io/gorm"
)

var db *gorm.DB

type Menu struct {
	gorm.Model
	Name        string `db:"name"`
	Code        string `db:"code"`
	Status      string `db:"status"`
	Category    string `db:"category"`
	Price       int32  `db:"price"`
	Description string `db:"description"`
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
	db := db.Where("CODE=?", Code).Find(&getMenu)
	return &getMenu, db
}

func (m *Menu) CreateMenu() *Menu {
	db.Create(m)
	return m
}

func (m *Menu) UpdateMenu(ID uint) *Menu {
	m.ID = ID
	db.Save(m)
	return m
}

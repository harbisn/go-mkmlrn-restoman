package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/pkg/database"
	"strings"
	"time"
)

var db *pg.DB

var (
	location, _ = time.LoadLocation("Asia/Jakarta")
)

type Menu struct {
	ID          uint64    `json:"id" pg:"id, pk"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Category    string    `json:"category"`
	Price       int32     `json:"price"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedBy   string    `json:"UpdatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func init() {
	db = database.Connect()
	err := db.Model((*Menu)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp: false,
	})
	if err != nil {
		var alreadyExists = strings.Contains(err.Error(), "already exists")
		if !alreadyExists {
			panic(err)
		}
	}
}

func (m *Menu) CreateMenu() *Menu {
	m.CreatedAt = time.Now().UTC().In(location)
	m.UpdatedAt = time.Now().UTC().In(location)
	_, err := db.Model(m).Insert()
	if err != nil {
		return nil
	}
	return m
}

func GetAllMenu() []Menu {
	var Menus []Menu
	err := db.Model(&Menus).Select()
	if err != nil {
		return nil
	}
	return Menus
}

func GetMenuById(ID uint64) (*Menu, *pg.DB) {
	var getMenu Menu
	err := db.Model(&getMenu).Where("id = ?", ID).Select()
	if err != nil {
		return nil, nil
	}
	return &getMenu, db
}

func (m *Menu) UpdateMenu(ID uint64) *Menu {
	m.ID = ID
	m.UpdatedAt = time.Now().UTC().In(location)
	_, err := db.Model(m).WherePK().Update()
	if err != nil {
		return nil
	}
	return m
}

func DeleteMenu(ID uint64) (*Menu, *pg.DB) {
	var deletedMenu Menu
	_, err := db.Model(&deletedMenu).Where("id = ?", ID).Delete()
	if err != nil {
		return nil, nil
	}
	return &deletedMenu, db
}

package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
)

type RoomRepository struct {
	DB *pg.DB
}

func NewRoomRepository(db *pg.DB) *RoomRepository {
	InitializeRoomTable(db)
	return &RoomRepository{
		DB: db,
	}
}

func InitializeRoomTable(db *pg.DB) {
	err := db.Model((*models.Room)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

func (r *RoomRepository) Create(m *models.Room) *models.Room {
	return Create(r.DB, m).(*models.Room)
}

func (r *RoomRepository) GetAll(p models.PageableDto) ([]models.Room, error) {
	var rooms []models.Room
	if err := GetAll(r.DB, &rooms, p); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomRepository) GetById(ID uint64) *models.Room {
	var room models.Room
	GetById(r.DB, &room, ID)
	return &room
}

func (r *RoomRepository) Update(m *models.Room) *models.Room {
	return Update(r.DB, m).(*models.Room)
}

func (r *RoomRepository) Delete(ID uint64) *models.Room {
	var room models.Room
	return Delete(r.DB, &room, ID).(*models.Room)
}

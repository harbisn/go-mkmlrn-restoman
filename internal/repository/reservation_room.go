package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
)

type ReservationRoomRepository struct {
	DB *pg.DB
}

func NewReservationRoomRepository(db *pg.DB) *ReservationRoomRepository {
	InitializeReservationRoomTable(db)
	return &ReservationRoomRepository{
		DB: db,
	}
}

func InitializeReservationRoomTable(db *pg.DB) {
	err := db.Model((*models.ReservationRoom)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

func (r *ReservationRoomRepository) Create(m *models.ReservationRoom) *models.ReservationRoom {
	return Create(r.DB, m).(*models.ReservationRoom)
}

func (r *ReservationRoomRepository) GetAll(p models.PageableDto) ([]models.ReservationRoom, error) {
	var ReservationRooms []models.ReservationRoom
	if err := GetAll(r.DB, &ReservationRooms, p); err != nil {
		return nil, err
	}
	return ReservationRooms, nil
}

func (r *ReservationRoomRepository) Update(m *models.ReservationRoom) *models.ReservationRoom {
	return Update(r.DB, m).(*models.ReservationRoom)
}

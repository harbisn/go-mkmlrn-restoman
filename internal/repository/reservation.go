package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
)

type ReservationRepository struct {
	DB *pg.DB
}

func NewReservationRepository(db *pg.DB) *ReservationRepository {
	InitializeReservationTable(db)
	return &ReservationRepository{
		DB: db,
	}
}

func InitializeReservationTable(db *pg.DB) {
	err := db.Model((*models.Reservation)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

func (r *ReservationRepository) Create(m *models.Reservation) *models.Reservation {
	return Create(r.DB, m).(*models.Reservation)
}

func (r *ReservationRepository) GetAll(p models.PageableDto) ([]models.Reservation, error) {
	var reservations []models.Reservation
	if err := GetAll(r.DB, &reservations, p); err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *ReservationRepository) GetById(ID uint64) *models.Reservation {
	var reservation models.Reservation
	GetById(r.DB, &reservation, ID)
	return &reservation
}

func (r *ReservationRepository) Update(m *models.Reservation) *models.Reservation {
	return Update(r.DB, m).(*models.Reservation)
}

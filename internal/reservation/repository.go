package reservation

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
	"time"
)

type Repository struct {
	DB *pg.DB
}

func NewReservationRepository(db *pg.DB) *Repository {
	initializeTable(db)
	return &Repository{
		DB: db,
	}
}

func initializeTable(db *pg.DB) {
	err := db.Model((*Reservation)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

func (r *Repository) Insert(reservation *Reservation) error {
	if _, err := r.DB.Model(reservation).Insert(); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Select(pageable pagination.PageableDto) ([]Reservation, int, error) {
	var reservations []Reservation
	query := pagination.SetFilterAndPagination(r.DB.Model(&reservations), pageable)
	var count int
	count, err := query.SelectAndCount()
	if err != nil {
		return nil, 0, err
	}
	return reservations, count, nil
}

func (r *Repository) SelectById(id uint64) (*Reservation, error) {
	var reservation Reservation
	if err := r.DB.Model(&reservation).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *Repository) Update(reservation *Reservation) error {
	if _, err := r.DB.Model(reservation).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindReservationsWithinTimeRange(roomId uint64, startAt, endAt time.Time) ([]Reservation, error) {
	var reservations []Reservation
	query := r.DB.Model(&reservations).Where("room_id = ?", roomId)
	query = query.Where("date(start_at) = ?", startAt)
	query = query.Where("? between start_at and end_at", startAt.Add(15*time.Minute))
	query = query.WhereOr("? between start_at and end_at", endAt.Add(-15*time.Minute))
	query = query.WhereOr("start_at between ? and ?", startAt, endAt)
	query = query.WhereOr("end_at between ? and ?", startAt, endAt)
	if err := query.Select(); err != nil {
		return nil, err
	}
	return reservations, nil
}

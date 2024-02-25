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
	var reservationsFromDB []Reservation
	query := pagination.SetFilterAndPagination(r.DB.Model(&reservationsFromDB), pageable)
	var count int
	count, err := query.SelectAndCount()
	if err != nil {
		return nil, 0, err
	}
	reservations := MapListToWithLocalTime(reservationsFromDB)
	return reservations, count, nil
}

func (r *Repository) SelectById(id uint64) (*Reservation, error) {
	var reservationFromDB Reservation
	if err := r.DB.Model(&reservationFromDB).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}
	reservation := MapToWithLocalTime(reservationFromDB)
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
	err := r.DB.Model(&reservations).Where("room_id = ?", roomId).
		Where("date(start_at) = ?", startAt).
		Where("? between start_at and end_at", startAt.Add(15*time.Minute)).
		WhereOr("? between start_at and end_at", endAt.Add(-15*time.Minute)).
		WhereOr("start_at between ? and ?", startAt, endAt).
		WhereOr("end_at between ? and ?", startAt, endAt).
		Select()
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *Repository) FindByDate(roomId uint64, startAt time.Time) ([]Reservation, error) {
	var reservationsFromDB []Reservation
	err := r.DB.Model(&reservationsFromDB).Where("room_id = ?", roomId).
		Where("date(start_at) = ?", startAt).Select()
	if err != nil {
		return nil, err
	}
	reservations := MapListToWithLocalTime(reservationsFromDB)
	return reservations, nil
}

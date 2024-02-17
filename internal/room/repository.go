package room

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
)

type Repository struct {
	DB *pg.DB
}

func NewRoomRepository(db *pg.DB) *Repository {
	initializeTable(db)
	return &Repository{
		DB: db,
	}
}

func initializeTable(db *pg.DB) {
	err := db.Model((*Room)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

func (r *Repository) Insert(room *Room) error {
	if _, err := r.DB.Model(room).Insert(); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Select(pageable pagination.PageableDto) ([]Room, error) {
	var rooms []Room
	query := pagination.SetFilterAndPagination(r.DB.Model(&rooms), pageable)
	if err := query.Select(); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *Repository) SelectById(id uint64) (*Room, error) {
	var room Room
	if err := r.DB.Model(&room).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *Repository) Update(room *Room) error {
	if _, err := r.DB.Model(room).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

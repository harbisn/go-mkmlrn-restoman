package menu

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/internal/helper/pagination"
)

type Repository struct {
	DB *pg.DB
}

func NewMenuRepository(db *pg.DB) *Repository {
	initializeTable(db)
	return &Repository{
		DB: db,
	}
}

func initializeTable(db *pg.DB) {
	err := db.Model((*Menu)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

func (r *Repository) Insert(menu *Menu) error {
	if _, err := r.DB.Model(menu).Insert(); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Select(pageable pagination.PageableDto) ([]Menu, int, error) {
	var menus []Menu
	query := pagination.SetFilterAndPagination(r.DB.Model(&menus), pageable)
	var count int
	count, err := query.SelectAndCount()
	if err != nil {
		return nil, 0, err
	}
	return menus, count, nil
}

func (r *Repository) SelectById(id uint64) (*Menu, error) {
	var menu Menu
	if err := r.DB.Model(&menu).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *Repository) Update(menu *Menu) error {
	if _, err := r.DB.Model(menu).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

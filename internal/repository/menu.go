package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
)

type MenuRepository struct {
	DB *pg.DB
}

func NewMenuRepository(db *pg.DB) *MenuRepository {
	InitializeMenuTable(db)
	return &MenuRepository{
		DB: db,
	}
}

func InitializeMenuTable(db *pg.DB) {
	err := db.Model((*models.Menu)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

func (r *MenuRepository) Create(m *models.Menu) *models.Menu {
	return Create(r.DB, m).(*models.Menu)
}

func (r *MenuRepository) GetAll(p models.PageableDto) ([]models.Menu, error) {
	var menus []models.Menu
	if err := GetAll(r.DB, &menus, p); err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *MenuRepository) GetById(ID uint64) *models.Menu {
	var menu models.Menu
	GetById(r.DB, &menu, ID)
	return &menu
}

func (r *MenuRepository) Update(m *models.Menu) *models.Menu {
	return Update(r.DB, m).(*models.Menu)
}

func (r *MenuRepository) Delete(ID uint64) *models.Menu {
	var menu models.Menu
	return Delete(r.DB, &menu, ID).(*models.Menu)
}

package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/harbisn/go-mkmlrn-restoman/internal/models"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var location, _ = time.LoadLocation("Asia/Jakarta")

func GetCurrentTime() time.Time {
	return time.Now().UTC().In(location)
}

func Create(db *pg.DB, model interface{}) interface{} {
	currentTime := GetCurrentTime()
	reflect.ValueOf(model).Elem().FieldByName("CreatedAt").Set(reflect.ValueOf(currentTime))
	reflect.ValueOf(model).Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(currentTime))
	_, err := db.Model(model).Insert()
	if err != nil {
		return nil
	}
	return model
}

func GetAll(db *pg.DB, model interface{}, p models.PageableDto) error {
	query := db.Model(model).Limit(p.Size).Offset(p.Offset)
	if p.Order != "" {
		sorters := strings.Split(p.Order, ",")
		for _, sorter := range sorters {
			query.Order(sorter)
		}
	}
	pattern := regexp.MustCompile(`^(lowest|highest)`)
	for key, value := range p.Filter {
		if strings.Contains(key, "highest") {
			key = pattern.ReplaceAllString(key, "")
			query = query.Where(key+" <= ?", value)
		} else if strings.Contains(key, "lowest") {
			key = pattern.ReplaceAllString(key, "")
			query = query.Where(key+" >= ?", value)
		} else {
			query = query.Where(key+" = ?", value)
		}
	}
	if err := query.Select(); err != nil {
		return err
	}
	return nil
}

func GetById(db *pg.DB, model interface{}, ID uint64) interface{} {
	err := db.Model(model).Where("id = ?", ID).Select()
	if err != nil {
		return nil
	}
	return model
}

func Update(db *pg.DB, model interface{}) interface{} {
	currentTime := GetCurrentTime()
	reflect.ValueOf(model).Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(currentTime))
	_, err := db.Model(model).WherePK().Update()
	if err != nil {
		return nil
	}
	return model
}

func Delete(db *pg.DB, model interface{}, ID uint64) interface{} {
	_, err := db.Model(model).Where("id = ?", ID).Delete()
	if err != nil {
		return nil
	}
	return model
}

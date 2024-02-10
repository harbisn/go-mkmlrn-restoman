package models

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/harbisn/go-mkmlrn-restoman/pkg/database"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var db *pg.DB

var (
	location, _ = time.LoadLocation("Asia/Jakarta")
)

func init() {
	db = database.Connect()
	InitializeTables((*Menu)(nil), (*Room)(nil), (*Reservation)(nil))
}

func InitializeTables(models ...interface{}) {
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			panic(err)
		}
	}
}

func GetCurrentTime() time.Time {
	return time.Now().UTC().In(location)
}

func Create(model interface{}) interface{} {
	currentTime := GetCurrentTime()
	reflect.ValueOf(model).Elem().FieldByName("CreatedAt").Set(reflect.ValueOf(currentTime))
	reflect.ValueOf(model).Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(currentTime))
	_, err := db.Model(model).Insert()
	if err != nil {
		return nil
	}
	return model
}

func GetAll(models interface{}, offset, size int, order string, filters map[string]interface{}) error {
	query := db.Model(models).Limit(size).Offset(offset)
	if order != "" {
		sorters := strings.Split(order, ",")
		for _, sorter := range sorters {
			query.Order(sorter)
		}
	}
	pattern := regexp.MustCompile(`^(lowest|highest)`)
	for key, value := range filters {
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

func GetById(model interface{}, ID uint64) interface{} {
	err := db.Model(model).Where("id = ?", ID).Select()
	if err != nil {
		return nil
	}
	return model
}

func Update(model interface{}) interface{} {
	currentTime := GetCurrentTime()
	reflect.ValueOf(model).Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(currentTime))
	_, err := db.Model(model).WherePK().Update()
	if err != nil {
		return nil
	}
	return model
}

func Delete(model interface{}, ID uint64) interface{} {
	_, err := db.Model(model).Where("id = ?", ID).Delete()
	if err != nil {
		return nil
	}
	return model
}

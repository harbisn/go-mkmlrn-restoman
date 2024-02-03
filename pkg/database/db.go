package database

import (
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=your_host user=your_user password=your_password dbname=your_dbname port=your_port sslmode=disable"
	d, err := gorm.Open(psql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

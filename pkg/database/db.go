package database

import (
	"github.com/spf13/viper"
	psql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := viper.GetString("database.dsn")
	d, err := gorm.Open(psql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	viper.SetEnvPrefix("GO_MKMLRN_RESTOMAN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func init() {
	loadConfig()
	requiredFields := []string{"database.dsn"}
	for _, field := range requiredFields {
		if !viper.IsSet(field) {
			log.Fatalf("Missing required configuration field: %s", field)
			return
		}
	}
}

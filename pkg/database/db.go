package database

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func Connect() (con *pg.DB) {
	options := &pg.Options{
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("database.host"), viper.GetString("database.port")),
		Database: viper.GetString("database.dbName"),
		PoolSize: viper.GetInt("database.poolSize"),
	}
	con = pg.Connect(options)
	if con == nil {
		log.Fatalf("cannot connect to postgres")
	}
	return
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
	requiredFields := []string{"database.user", "database.password", "database.host",
		"database.port", "database.dbName", "database.poolSize"}

	for _, field := range requiredFields {
		if !viper.IsSet(field) {
			log.Fatalf("Missing required configuration field: %s", field)
			return
		}
	}

}

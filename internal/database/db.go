package database

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func Connect() (db *pg.DB) {
	loadConfig()
	options := &pg.Options{
		User:     viper.GetString("datasource.user"),
		Password: viper.GetString("datasource.password"),
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("datasource.host"), viper.GetString("datasource.port")),
		Database: viper.GetString("datasource.database"),
		PoolSize: viper.GetInt("datasource.poolSize"),
	}
	db = pg.Connect(options)
	if db == nil {
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

	requiredFields := []string{"datasource.user", "datasource.password", "datasource.host",
		"datasource.port", "datasource.database", "datasource.poolSize"}

	for _, field := range requiredFields {
		if !viper.IsSet(field) {
			log.Fatalf("Missing required configuration field: %s", field)
			return
		}
	}

}

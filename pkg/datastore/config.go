package datastore

import (
	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

func dbConfig() *pg.Options {
	host := viper.GetString("db_host")
	user := viper.GetString("db_username")
	password := viper.GetString("db_password")
	name := viper.GetString("db_name")

	options := pg.Options{
		Addr:     host,
		User:     user,
		Password: password,
		Database: name,
	}

	return &options
}

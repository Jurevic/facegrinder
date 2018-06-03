package datastore

import (
	"fmt"
	"github.com/go-pg/pg"
)

var db *pg.DB

func Init() {
	options := dbConfig()

	db = pg.Connect(options)

	fmt.Println("Successfully connected!")
}

func Close() {
	db.Close()
}

func DBConn() *pg.DB {
	return db
}

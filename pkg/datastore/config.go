package datastore

import (
	"os"
	"github.com/go-pg/pg"
)


const (
	dbaddr = "DBADDR"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func dbConfig() *pg.Options {
	addr, ok := os.LookupEnv(dbaddr)
	if !ok {
		panic("DBADDR environment variable required but not set")
	}

	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}

	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}

	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}

	options := pg.Options{
		Addr:addr,
		User:user,
		Password:password,
		Database:name,
	}

	return &options
}

package migrate

import (
	"fmt"

	"github.com/go-pg/migrations"
)

const processorsTable = `
CREATE TABLE processors (
id serial NOT NULL,
owner_id integer NOT NULL,
name text NOT NULL,
nodes jsonb,
PRIMARY KEY (id),
FOREIGN KEY (owner_id) REFERENCES users(id)
)`

const facesTable = `
CREATE TABLE faces (
id serial NOT NULL,
owner_id integer NOT NULL,
path text NOT NULL,
name text NOT NULL,
PRIMARY KEY (id),
FOREIGN KEY (owner_id) REFERENCES users(id)
)`

const usersTable = `
CREATE TABLE users (
id serial NOT NULL,
password bytea NOT NULL,
email text NOT NULL UNIQUE,
name text NOT NULL,
is_active boolean NOT NULL DEFAULT TRUE,
is_superuser boolean NOT NULL DEFAULT FALSE,
created_at timestamp NOT NULL DEFAULT current_timestamp,
updated_at timestamp,
last_login timestamp,
PRIMARY KEY (id)
)`

func init() {
	up := []string{
		usersTable,
		processorsTable,
		facesTable,
	}

	down := []string{
		`DROP TABLE processors`,
		`DROP TABLE faces`,
		`DROP TABLE users`,
	}

	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating initial tables")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping initial tables")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

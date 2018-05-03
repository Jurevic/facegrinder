package migrate

import (
	"fmt"

	"github.com/go-pg/migrations"
)

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
	}

	down := []string{
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

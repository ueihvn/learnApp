package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Connect("postgres", "user=admin dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

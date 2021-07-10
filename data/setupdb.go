package data

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	Db                *sqlx.DB
)

func init() {
	LoadEnv()

	POSTGRES_USER = getEnvByKey("POSTGRES_USER")
	POSTGRES_PASSWORD = getEnvByKey("POSTGRES_PASSWORD")
	POSTGRES_DB = getEnvByKey("POSTGRES_DB")

	connectionString := "user=" + POSTGRES_USER + " dbname=" + POSTGRES_DB + " password=" + POSTGRES_PASSWORD + " sslmode=disable"
	fmt.Println(connectionString)

	var err error
	Db, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=admin dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		fmt.Println(Db)
	}

}

func main() {
	fmt.Println("testing connect to postgresdb")
}

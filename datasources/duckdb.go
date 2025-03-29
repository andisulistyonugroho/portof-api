package datasources

import (
	"database/sql"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

func Duckconnect() *sql.DB {

	// fmt.Println("BEGINING DS: ", DB)

	DB, err := sql.Open("duckdb", "../portof.ddb?access_mode=read_only")

	if err != nil {
		log.Fatal("Failed to connect duckdb database..." + err.Error())
	}

	return DB
}

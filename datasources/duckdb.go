package datasources

import (
	"database/sql"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

// var DB *sql.DB

func Duckconnect() (*sql.DB) {

	// fmt.Println("BEGINING DS: ", DB)

	DB, err := sql.Open("duckdb","../portof.ddb")
	
	if err != nil {
		log.Fatal("Failed to connect duckdb database..." + err.Error())
	}
	defer DB.Close()

	return DB

	// var (
	// 	ID int
	// 	Username string
	// )

	// row := DB.QueryRow("SELECT id,username FROM users")
	// fmt.Println(row)
	// erro := row.Scan(&ID,&Username)
	// if erro != nil {
	// 	fmt.Println(erro.Error())
	// 	return
	// }

	// fmt.Printf("id: %d, name: %s\n", ID, Username)
}
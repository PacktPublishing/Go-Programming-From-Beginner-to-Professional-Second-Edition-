package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var name string
	var id int
	id = 2
	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}
	qryrow, err := db.Prepare("SELECT name FROM test WHERE id=$1")
	if err != nil {
		panic(err)
	}
	err = qryrow.QueryRow(id).Scan(&name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The name with id %d is %s", id, name)
	err = qryrow.Close()
	if err != nil {
		panic(err)
	}
	db.Close()
}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID          int
	FirstName   string
	LastName    string
	NameOfGroup string
	Status      bool
}

func main() {
	connStr := "user=postgres dbname=firstDB sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error while connecting")
		panic(err)
	}
	defer db.Close()
	result := db.QueryRow(`select firstname, lastname FROM firsttable
   WHERE ID = $1`, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	//fmt.Println(result.RowsAffected())
}

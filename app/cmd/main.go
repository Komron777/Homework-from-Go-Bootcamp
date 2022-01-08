/*
We have a database:

id | firstname | lastname | nameofgroup | status 
----+-----------+----------+-------------+--------
  2 | Elon      | Musk     | C--         | f
 10 | Bill      | Gates    | ,NET        | t

The task is to change the inforamtion in the given database in order to create functions that are 
connected to interfaces which are also connected to structs which have methods. 
Just a crazy explanation! :)
*/

package main

import (
	f "app/funcs"
	"database/sql"
	//i "app/interfaces" //--> put this package in comment unless you print output
	//"fmt" // --> as well as "fmt" package 
	_ "github.com/lib/pq"

)

var (
	connStr = "user=komron dbname=users password=esdmk sslmode=disable"
)

func main() {
	db, _ := sql.Open("postgres", connStr)
	f.DB = db
	defer db.Close()


	//READY to print!
	//structPtr := &f.User{}
	// forUPDATE := f.User{ID: 10, FirstName: "Jahongir", LastName: "Anorboyev", NameOfGroup: "Go", Status: true}
	//forINSERT := f.User{ID: 10,FirstName: "Bill", LastName: "Gates", NameOfGroup: ",NET", Status: true}
	// forDELETE := 3
	// forGETBYID := 2
	// fmt.Println(i.GETBYID(structPtr,forGETBYID))
	//fmt.Println(i.GET(structPtr))
	//fmt.Println(i.INSERT(structPtr,forINSERT))
	// fmt.Println(i.UPDATE(structPtr,forUPDATE))
	// fmt.Println(i.DELETE(structPtr,forDELETE))	
}


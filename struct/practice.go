package main

import "fmt"

type User struct {
	fName string
	lName string
}

// func concatanation(u User) string {
// 	return u.fName + u.lName
// }
// func main() {
// 	users := []User{
// 		{
// 			fName: "Komron",
// 			lName: "Fayziboyev",
// 		},
// 		{
// 			fName: "Davron",
// 			lName: "Yo'lchiboyev",
// 		},
// 	}
// 	fmt.Println(concatanation(users[0]))
// }

//func (u User) Concatination() string -- METHOD , dependent on *struct

type Son struct {
	number int
}
//
func (s Son) read()  {
	birlik := []string {"","bir","ikki","uch","to'rt","besh","olti","yetti","sakkiz","to'qqiz"}
	onlik := []string {"","o'n","yigirma","o'ttiz","qirq","ellik","oltmish","yetmish","sakson","to'qson"}
	fmt.Println(onlik[s.number/10],birlik[s.number%10])
}

func main() {
	var son Son = Son{number: 26}
	son.read()
}

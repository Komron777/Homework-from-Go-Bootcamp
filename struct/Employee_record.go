package main

import (
	"fmt"
)

type User struct {
	ID     uint
	FName  string
	LName  string
	Age    int
	Status bool
}

func (User)Update(users []User) []User{
	u1 := User {
		ID: 0,
		FName: "Khurshidbek",
		LName: "Sobirov",
		Age: 20,
		Status: false, 
	}
	for i,_ := range users {
		if users[i].ID == u1.ID {
			users[i].FName = u1.FName
			users[i].LName = u1.LName
		}
	}
	return users 
}

func main() {
	users := []User{
		{
			ID: 01,
			FName: "Komron",
			LName: "Fayziboyev",
			Age: 17,
			Status: true, 
		},
		{
			ID: 02,
			FName: "Javohir",
			LName: "Xoliqberdiyev",
			Age: 17,
			Status: false, 
		},
		{
			ID: 03,
			FName: "Elyor",
			LName: "Tukhtamurodov",
			Age: 17,
			Status: true, 
		},
		{
			ID: 04,
			FName: "Davron",
			LName: "Yo'lchiboyev",
			Age: 25,
			Status: false, 
		},
	}
	fmt.Print("The existing information:\n",users)
	fmt.Println(users)
}
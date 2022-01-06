package main

import (
	"encoding/json"
	"fmt"
)
type mainStruct struct {
	Users []User `json:"users"`
}
type User struct {
	FName string `json:"first_name"`
	LName string `json:"last_name"`
	Score int    `json:"score"`
}

var json1 = `
{
	"users" : [
		{
			"first_name" : "Komron",
			"last_name" : "Fayziboev",
			"score" : 23
		},
		{
			"first_name" : "Akmal",
			"last_name" : "Akhadjonov",
			"score" : 142
		},
		{
			"first_name" : "Bakhtishod",
			"last_name" : "Umurzakov",
			"score" : 22
		}
	]
}
`

func main() {
	var sum int
	p := mainStruct{}
	if err := json.Unmarshal([]byte(json1), &p); err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(p.Users); i++{
		sum += p.Users[i].Score
	}
	avg := sum/len(p.Users)
	fmt.Println(avg)
}

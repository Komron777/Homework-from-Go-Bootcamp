package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	Name string `json:"name"`
	Age int `json:",omitempty"`
	Status bool `json:"-"`
}
func main() {
	m := myStruct{"Komron",0,true}
	fmt.Println(m)
	data,err := json.Marshal(m)
	if err != nil {
		fmt.Println(data)
		return
	}
	fmt.Printf("%s\n",data)
}
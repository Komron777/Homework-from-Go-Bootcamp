package main

import (
	"encoding/json"
	"fmt"
)
type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}
func main() {
s := myStruct{
	Name:   "John Connor",
	Age:    35,
	Status: true,
	Values: []int{15, 11, 37},
}

// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
// и возвращает байтовый срез с данными, кодированными в формат JSON.
data, err := json.Marshal(s)
if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("%s", data) // {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
for _,j := range data {
	fmt.Println(string(j))	
}
}
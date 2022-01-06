/*
https://stepik.org/lesson/353243/step/6?unit=337227

if JSON file is local
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type mainStruct struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}
type Student struct {
	LastName    string
	FirstName   string
	MiddleName  string
	DateOfBirth string
	Address     string
	Phone       string
	Rating      []int
}
type Average struct {
	Average float64
}

var infoAboutGroup = `{
	"ID" : 134,
	"Number" : "AF-1",
	"Year" : 2,
	"Students" : [
		{
			"LastName" : "Fayziboyev",
			"FirstName" : "Komron",
			"MiddleName" : "Erkinjon o'g'li",
			"DateOfBirth" : "15.01.2004",
			"Address" : "Furkat 7, Gazalkent city, Bostonlyk dis., Tashkent reg.",
			"Phone" : "+998977777777",
			"Rating": [23,5,6,7,4,7]
		},
		{
			"LastName" : "Akhadjonov",
			"FirstName" : "Akmal",
			"MiddleName" : "Asrorjon o'g'li",
			"DateOfBirth" : "22.04.2004",
			"Address" : "Fergana",
			"Phone" : "+998977347777",
			"Rating": [6,7,4]
		},
		{
			"LastName" : "Umurzakov",
			"FirstName" : "Bakhtishod",
			"MiddleName" : "Eshnazar o'g'li",
			"DateOfBirth" : "22.11.2004",
			"Address" : "Qashqadaryo",
			"Phone" : "+998977771111",
			"Rating": [23,4,5,456,7,9,5,4,7,8]
		}
	]
}`

func main() {
	var allGrades float64
	const fileName = "info.json"
	ByteVersion := []byte(infoAboutGroup)
	file, err := os.Create(fileName)
	
	defer file.Close()
	if err := ioutil.WriteFile(fileName, ByteVersion, 0666); err != nil {
		fmt.Println(err)
	}
	///////////
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	s := mainStruct{}
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(s.Students); i++ {
		allGrades += float64(len(s.Students[i].Rating))
	}
	avr := Average{
		Average: allGrades / float64(len(s.Students)),
	}
	d, err := json.MarshalIndent(avr, "", "    ")
	fmt.Println(string(d))
}

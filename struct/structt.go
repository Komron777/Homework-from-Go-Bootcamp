package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Siblings []Sibling
}
type Sibling struct {
	Name string
	Age  int
	Type string
}

func main() {
	persons := []Person{
		{
			Name: "Akmal",
			Age:  17,
			Siblings: []Sibling{
				{
					Name: "Akbarshox",
					Age:  12,
					Type: "little brother",
				},
				{
					Name: "Komron",
					Age:  17,
					Type: "friend",
				},
			},
		},
	}
	fmt.Println(persons)
}

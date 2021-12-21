package main

import "fmt"

type EduCenter struct {
	Name    string
	Year    int
	Boss    string
	Branchs []Branch
}
type Branch struct {
	Adress string
	Year   int
	Tutors []Tutor
}
type Tutor struct {
	Name  string
	Level string
}

func main() {

	centers := []EduCenter{
		{
			Name: "NT",
			Year: 2018,
			Boss: "TA",
			Branchs: []Branch{
				{
					Adress: "",
					Year:   2018,
					Tutors: []Tutor{
						{
							Name:  "",
							Level: "",
						},
					},
				},
				{
					Adress: "",
					Year:   2018,
					Tutors: []Tutor{
						{
							Name:  "",
							Level: "",
						},
					},
				},
				{
					Adress: "",
					Year:   2018,
					Tutors: []Tutor{
						{
							Name:  "",
							Level: "",
						},
					},
				},
			},
		},
		{
			"NT2",
			2018,
			"TA2",
			[]Branch{
				{
					Adress: "2",
					Year:   2018,
					Tutors: []Tutor{
						{
							"T2",
							"L2",
						},
						{
							Name:  "T2",
							Level: "L2",
						},
						{
							Name:  "T2",
							Level: "L2",
						},
						{
							Name:  "T2",
							Level: "L2",
						},
					},
				},
			},
		},
	}
	fmt.Println(centers)

	
}

package main

import "fmt"

type Instrument interface{
	play()
}
type Drum struct {
	size float32
}
type Guitar struct {
	numberOfStrings int
}
type Trumpet struct {
	diameter float32
}

func (d Drum) play() {
	fmt.Printf("Drum which has a size of %v is playing\n",d.size)
} 
func (g Guitar) play() {
	fmt.Printf("Guitar which has %v stirngs is playing\n",g.numberOfStrings)
} 
func (t Trumpet) play() {
	fmt.Printf("Trumpet which has a diameter of %v is playing",t.diameter)
} 

func main() {
	ins1 := Drum{465.6}
	ins2 := Guitar{14}
	ins3 := Trumpet{0.2}
	instruments := [...]Instrument{ins1,ins2,ins3}
	for _,v := range instruments {
		v.play()
	}
}
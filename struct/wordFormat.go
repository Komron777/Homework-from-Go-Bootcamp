package main

import (
	"fmt"
)

type Word struct {
	word string
}

func (w *Word) Upper () string {
	var text1 string
	for _,j := range w.word {
		if j >= 97 && j <= 122 {
			text1 += string(j-32)
		} else {
			text1 += string(j)
		}
	}
	return text1
}

func (w *Word) Lower() string {
	var text1 string
	for _,j := range w.word {
		if j >= 65 && j <= 90 {
			text1 += string(j+32)
		} else {
			text1 += string(j)
		}
	}
	return text1
}

func (w * Word)Reverse() string {
	var text1 string
	for _,j := range w.word {
		if j >= 97 && j <= 122 {
			text1 += string(j-32)
		} else if j >= 65 && j <= 90 {
			text1 += string(j+32)
		} else {
			text1 += string(j)
		}
	}
	return text1
}

func main(){
	text := Word{"Private educational institutions always outperform the public ones"}
	fmt.Println(text.Upper())
	fmt.Println(text.Lower())
	fmt.Println(text.Reverse())
}

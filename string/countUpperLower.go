package main

import (
	"fmt"
)

func CountUpperLower(text string) (int,int) {
	var upper,lower int
	for _,j := range text {
		if j >= 65 && j <= 90 {
			upper++
		} else if j >= 97 && j <= 122 {
			lower++
		}
	}
	return upper,lower
}
func main(){
	fmt.Println(CountUpperLower("Hello UEFA"))

}
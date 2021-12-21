package main

import (
	"fmt"
	"strings"
)

func LastIndex(text string, letter string) (int) {
	n := strings.Count(text,letter)
	var count int
	for i,j := range text {
		if string(j) == letter {
			count++
		}
		if count == n {
			return i
		}
	}
	return 0
}
func main(){
	fmt.Println(LastIndex("Hello UEFA","l"))
}
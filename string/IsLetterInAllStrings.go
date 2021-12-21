package main

import (
	"fmt"
)

func CommonLettersInAllStrings(text1, text2 string) []string {
	var letters []string
	for _, v1 := range text1 {
		for _, v2 := range text2 {
			if v1 == v2 && check(string(v1), letters) && v1 != 32 {
				letters = append(letters, string(v1))
			}
		}
	}
	return letters
}
func check(letter string, slice []string) bool {
	for _, v := range slice {
		if string(v) == letter {
			return false
		}
	}
	return true

}
func main() {
	fmt.Println(CommonLettersInAllStrings("Hello UEFA", "Portugal football team"))
}

package main

import (
	"fmt"
	"strings"
	//"strings"
)

func main() {
	// s := "Hello world" + "wassup?"
	// s1 := `Salom dunyo`

	// for _,val := range s {
	// 	fmt.Println(string(val))

	// }
	// fmt.Println(strings.Contains("Komron","om"))//boolean
	// fmt.Println(strings.Count("Komron","o"))//boolean
	// fmt.Println(strings.HasPrefix("Komron","Ko"))//boolean
	// fmt.Println(strings.HasSuffix("Komron","on"))//boolean
	// fmt.Println(strings.Index("Komron","Kosdfd"))//boolean
	// fmt.Println(strings.Join(["Komron","Fayziboyev","Komron"] []string,"**")//boolean
	// fmt.Println(strings.Index("Komron","Kosdfd"))//boolean
	// fmt.Println(strings.Trim("Komron","or"))
	// var word, word1 string
	// fmt.Scanln(&word)
	// for _,j := range word{
	// 	word1 = string(j) + word1
	// }
	// if word == word1 {
	// 	fmt.Println(true)
	// } else {
	// 	fmt.Println(false)	
	// }
	text := "Hello Tashkent!"
	var new string
	for _,j := range text {
		new += string(j)
	}
}

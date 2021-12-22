package main

import (
	f "fmt"
	s "strings"
//	u "unicode"
	//"strings"
)

func main() {
	// s := "Hello world" + "wassup?"
	// s1 := `Salom dunyo`

	// for _,val := range s {
	// 	fmt.Println(string(val))

	// }
	// fmt.Println(strings.Contains("Komron","om"))//boolean
	// fmt.Println(strings.Count("Komron","o"))
	// fmt.Println(strings.HasPrefix("Komron","Ko"))
	// fmt.Println(strings.HasSuffix("Komron","on"))
	// fmt.Println(strings.Index("Komron","Kosdfd"))
	// fmt.Println(strings.Join(["Komron","Fayziboyev","Komron"] []string,"**")
	// fmt.Println(strings.Index("Komron","Kosdfd"))
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
	f.Println(s.Join([]string{"Komron","Akmal","Bakhtishod"}," is a IHT student\n"))
}

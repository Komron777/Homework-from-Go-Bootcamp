package main
import (
	"fmt"
	"unicode/utf8"
)
func main(){
	var en = "english"
	var ru = "русский"
	fmt.Println(len(en),len(ru))
	fmt.Println(utf8.RuneCountInString(en),utf8.RuneCountInString(ru))
}
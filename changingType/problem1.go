package main

import (
	//"bytes"
	"fmt"
	"strconv"
	//"strings"
	"unicode"
)
func addNumbers(a,b string) int64 {
	var n1,n2 int64
	var num string
	for _,j := range a {
		if unicode.IsDigit(j) {
			num += string(j)
		}
	}
	d,_ := strconv.ParseInt(string(num),10,64)
	n1 = d
	num = ""
	for _,j := range b {
		if unicode.IsDigit(j) {
			num += string(j)
		}
	}
	d,_ = strconv.ParseInt(string(num),10,64)
	n2 = d
	return n1 + n2
}
func main () {
	fmt.Println(addNumbers("%^80","hhhhh20&&&&nd"))
}
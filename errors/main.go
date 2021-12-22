package main

import (
	"fmt"
	//"strconv"
	//"error"
)
func main(){
	var a uint32 = 5
	defer specPrint(a)
	a = 6
	fmt.Print(a)
}
func specPrint(s uint32) uint32 {
	fmt.Print(s)
	s++
	return s
}
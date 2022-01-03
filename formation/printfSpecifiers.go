package main

import (
	"fmt"
)
func main() {
	a := true
	b := 244654999
	c := 'й'
	d := 62.3386989
	var e float32 = 34
	fmt.Printf("%t\n",a) //boolean
	fmt.Printf("%b\n",b) //binary
	fmt.Printf("%c\n",c) //rune
	fmt.Printf("%d\n",b) //decimal
	fmt.Printf("%o\n",b) //octal
	fmt.Printf("%q\n",c) //quote
	fmt.Printf("%x\n",b) //hexagon
	fmt.Printf("%X\n",b) //hexagon
	fmt.Printf("%U\n",c) //unicode
	fmt.Printf("%e\n",d) //float64
	fmt.Printf("%f\n",e) //float64 & float32
	fmt.Printf("%g\n",d) //float64
//	fmt.Printf("%p\n",a) //pointer
	fmt.Printf("%.f\n",d) // to get rid fo fractional part of number
}
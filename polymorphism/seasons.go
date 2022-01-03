package main

import (
	"fmt"
)


const (
	a = iota + 1
	_
	b
	c
	d = c + 2
	i
	j
	i2 = iota + 2
	i3 
	i4

)

func main() {
	fmt.Println(b,c,d,t,i,i2,i3,i4)
}
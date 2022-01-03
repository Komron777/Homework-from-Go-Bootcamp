package main

import "fmt"

func main() {
	var a float64 = 123.544
	result := fmt.Sprintf("%.2f\n",a) // always returns data in type string
	fmt.Printf("%v",result)
}
package main

import "fmt"

func addition(numbers ...int) (a, b int) {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	a = len(numbers)
	b = sum
	return a, b
}
func main() {
	a, b := addition(1, 2, 3, 5, 78, 989, 12)
	fmt.Println(a, b)
}

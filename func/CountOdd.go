package main

import "fmt"

func OddCount(n int) int {
	var odd [] int
	for i := 1; i <= n; i += 2 {
		odd = append(odd, i)
	}
	return 0
}
func main() {
	fmt.Println(OddCount(23))
}

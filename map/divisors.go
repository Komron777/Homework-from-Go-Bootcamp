package main

import "fmt"

func showDivisors(n1, n2 int) map[int][]int {
	var start, finish int
	divisorsMap := make(map[int][]int)
	divisorSlice := []int{}
	if n1 > n2 {
		start, finish = n2, n1
	} else if n1 < n2 {
		start, finish = n1, n2
	}
	for i := start; i <= finish; i++ {
		divisorSlice = []int{}
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				divisorSlice = append(divisorSlice, j)
			}
		}
		divisorsMap[i] = divisorSlice
	}
	return divisorsMap
}
func main() {
	fmt.Println(showDivisors(12, 15))
}

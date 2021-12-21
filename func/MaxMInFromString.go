package main

import (
	"fmt"
	"strconv"
)

func MaxMin(numbers string) (max, min int) {
	var num string
	var integers []int
	for _, j := range numbers {
		if j != 32 {
			num += string(j)
		} else {
			digit, _ := strconv.Atoi(num)
			integers = append(integers, digit)
			num = ""
		}
	}
	max = integers[0]
	min = integers[0]
	for _, v := range integers {
		if max > v {
			max = v
		}
		if min < v {
			min = v
		}
	}
	return max, min
}
func main() {
	var numbers string = "4 7 9 0 68 0 78 6 4 5 6 87 8 -2 2"
	a, b := MaxMin(numbers)
	fmt.Printf("Max number: %v\nMin number: %v\n", a, b)
}

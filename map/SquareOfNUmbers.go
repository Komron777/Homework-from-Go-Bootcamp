package main

import "fmt"

func SliceToMap(a []int, b []int) map[int]int {
	//var k bool
	numAndSqrs := make(map[int]int)
	for _,v1 := range a{
		//k = false
		sqr := v1 * v1
		for _,v2 := range b {
			if sqr == v2 {
				numAndSqrs[v1] = v2
				//k = true
				break
			} else {
				numAndSqrs[v1] = 0
			}
		}
	}
	return numAndSqrs
}
func main() {
	a := []int{3,6,1,2,7,3,6,8}
	b := []int{24,36,9,1,64,23,4}
	fmt.Println(SliceToMap(a,b))
}

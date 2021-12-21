//recursion
package main

import "fmt"

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	noon := 0
	currentHeight := 0
	for {
		currentHeight += upSpeed
		noon++
		if currentHeight >= desiredHeight {
			break
		}
		currentHeight -= downSpeed
	}
	return noon
}

func NbDig(n int, d int) int {
	count := 0
	for i := 0; i <= n; i++ {
		s := i*i
		fmt.Print(s," ")
		for s != 0 {
			digit := s % 10
			if digit == d{
				count++
			}
			s /= 10
		}
	}
	if d == 0{
		return count + 1
	}else{
		return count
	}	
}
func main() {
	fmt.Println(NbDig(11,0))
}

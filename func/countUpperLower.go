package main

import "fmt"

func count(words []string) (upper, lower int) {
	for _, v := range words {
		for _, j := range v {
			if j >= 65 && j <= 90 {
				upper++
			} else if j >= 97 && j <= 122 {
				lower++
			}
		}
	}
	return upper, lower
}

func main() {
	var text = []string{"hello", "WoRld", "UEFA", "8/0 = error"}
	fmt.Println(count(text))
}

/*Another solution*/
// func main() {
// 	h := []string{"saLom", "QALE", "UEFA"}
// 	a, b := sortt(h...)
// 	fmt.Println("kichik: ", a)
// 	fmt.Println("katta: ", b)
// }
// func sortt(z...string) (a, b int) {
// 	for _, i := range z {
// 		for _, j := range i {
// 			if j >= 65 && j <= 90 {
// 				b++
// 			} else if j >= 97 && j <= 122 {
// 				a++
// 			}

// 		}
// 	}
// 	return a, b
// }

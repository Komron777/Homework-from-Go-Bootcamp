package main

//"strconv"

// func addition(numbers ...int)(a,b int){
// 	var sum int
// 	for _,v := range numbers {
// 		sum += v
// 	}
// 	a = len(numbers)
// 	b = sum
// 	return a,b
// }
// func main(){
// 	a, b := addition(1, 2, 3, 5,78,989, 12)
// 	fmt.Println(a, b)
// }

////////////////////

// func IsUpperCase(text string)(t bool){
// 	for _,j := range text {
// 		if !(j >= 65 && j <= 90) {
// 			return false
// 		}
// 	}
// 	return true
// }
// func main(){
// 	var text string
// 	fmt.Scanln(&text)
// 	fmt.Println(IsUpperCase(text))
// }
//salom
// func count(words []string)(upper,lower int){
// 	for _,v := range words {
// 		for _,j := range v {
// 			if j >= 65 && j <= 90 {
// 				upper++
// 			} else if j >= 97 && j <= 122 {
// 				lower++
// 			}
// 		}
// 	}
// 	return upper,lower
// }

// func main(){
// 	var text = []string{"hello","WoRld","UEFA","8/0 = error"}
// 	fmt.Println(count(text))
// }

///////////////////////////

// func MaxMin(numbers string) (max, min int) {
// 	var num string
// 	var integers []int
// 	for _, j := range numbers {
// 		if j != 32 {
// 			num += string(j)
// 		} else {
// 			digit, _ := strconv.Atoi(num)
// 			integers = append(integers, digit)
// 			num = ""
// 		}
// 	}
// 	max = integers[0]
// 	min = integers[0]
// 	for _,v := range integers {
// 		if max > v {
// 			max = v
// 		}
// 		if min < v {
// 			min = v
// 		}
// 	}
// 	return max, min
// }
// func main() {
// 	var numbers string = "4 7 9 0 68 0 78 6 4 5 6 87 8 -2 2"
// 	a,b := MaxMin(numbers)
// 	fmt.Printf("Max number: %v\nMin number: %v\n",a,b)
// }

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

//Khusniddin's version
// func main() {
// 	value := "salom. dunyo"
// 	value2 := ""
// 	for i, j := range value {
// 		if (i != 0 && value[i-1] == ' ' && value[i-2] == '.') && (j-32 >= 65 && j-32 <= 90) {
// 			value2 += string(j - 32)
// 		} else {
// 			value2 += string(j)
// 		}
// 	}
// 	value = value2
// 	fmt.Println(value)
// }

//hvghvgfhbvghvh

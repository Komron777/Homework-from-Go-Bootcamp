package main

import "fmt"

// type calculator interface {
// 	add(float64, float64) float64
// 	subtract(float64, float64) float64
// 	multi(float64, float64) float64
// 	div(float64, float64) float64
// }

// type adder interface {
// 	add(float64, float64) float64
// }

// type adderMulter interface {
// 	add(float64, float64) float64
// 	multi(float64, float64) float64
// }

// type myCalc struct{}

// func (m myCalc) add(a, b float64) float64 {
// 	return a + b
// }

// func (m myCalc) subtract(a, b float32) float64 {
// 	return float64(a - b)
// }

// func (m myCalc) multi(a, b float64) float64 {
// 	return float64(a * b)
// }

// func (m myCalc) div(a, b float64) float64 {
// 	return a / b
// }
// func main() {
// 	var i interface{} = 12

// 	if v, ok := i.(string); !ok {
// 		fmt.Println(v+"12") // Суммирование не произойдет, если ok == false
// 	}	
// }

////////////////

// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println("Умножим на 2:", v*2)
// 		fmt.Printf("%T\n",v)
// 	case string:
// 		fmt.Println(v + " golang")
// 	default:
// 		fmt.Printf("Я не знаю такого типа %T!\n", v)
// 	}
// }

// func main() {
// 	do(21)
// 	do("hello")
// 	do(true)
// }

/////////////////////////////
func readTask()(value1, value2, operation interface{}){
	return 4.7,2.5,"+"
}
// func printError(value interface{}){
// 	fmt.Printf("value = %v: %T\n",value,value)
//}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	v1, ok := value1.(float64) 
	if !ok {
		panic(value1)
	}
	if v, ok := operation.(string); ok {
		switch v {
			case "+"
		}
	}
}
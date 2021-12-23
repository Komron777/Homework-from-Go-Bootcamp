package main

import (
	"errors"
	"fmt"
	//"strconv"
	//"error"
)
func divide(a,b int) (int,error) {
	if b == 0 {
		return -1,errors.New("Can't be divided by zero!")
	}
	return a/b,nil
	//errors.Is()
}
func finish() {
	fmt.Println("Good bye")
}
func main(){
	//fmt.Println(divide(5,0))
	fmt.Println("Wassup")
	defer fmt.Println("bro")
	defer finish()
	fmt.Println("How're you doing?")
}
/*
In the first step, 10 positive integers are fed to standard input,
which must be written in the order they were entered into an array of 10 elements.
The type of numbers included in the array must correspond to the smallest possible unsigned integer.
The name of the array that you must create yourself workArray (the condition is required).
The fmt package has already been imported to read from standard input.

At the second stage, 3 more pairs of numbers are fed to the standard input -
the indices of the elements of this array, which need to be swapped
(if such a pair of numbers is 3 and 7, then in the array the element with index 3 must be swapped with the element whose index is 7).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n,a,b uint8
	var workArray [10] uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		workArray[i] = n
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&a,&b)
		workArray[a],workArray[b] = workArray[b],workArray[a]
	}
	for _,v := range workArray {
		fmt.Printf("%v ",v)
	
	k := bufio.NewReader(os.Stdin)
	fmt.Println(k)
	}
}
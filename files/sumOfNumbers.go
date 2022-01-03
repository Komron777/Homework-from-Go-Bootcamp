package main

import (
	"bufio"
	"fmt"
//	"io"
	"os"
	"strconv"
)

func main() {
	sum := 0
	file, err := os.Open("numbers.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		num,err := strconv.Atoi(s.Text())
		if err != nil{
			panic(err)
		}
		sum += num
	}
	os.Stdin()
}
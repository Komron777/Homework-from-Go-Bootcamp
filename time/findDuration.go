/*
Find the duration between two periods

Important: dates are enterd in string type as seperated by comma
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	timeArr := strings.Split(text, ",")
	//fmt.Println(timeArr[0])
	//fmt.Println(timeArr[1])

	t1, _ := time.Parse("02.01.2006 15:04:05", timeArr[0])
	t2, _ := time.Parse("02.01.2006 15:04:05", timeArr[1])
	if t1.Before(t2) {
		fmt.Println(t2.Sub(t1))
	} else {
		fmt.Println(t1.Sub(t2))
	}
}

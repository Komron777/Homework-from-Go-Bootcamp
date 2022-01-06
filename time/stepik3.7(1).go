/*
https://stepik.org/lesson/359395/step/3?unit=343626
*/
package main

import (
	"fmt"
	//"io/ioutil"
	//"os"
	"time"
)

func main() {
	//data,_ := ioutil.ReadAll(os.Stdin)
	date := "1986-04-16T05:20:00+06:00"
	t,_ := time.Parse(time.RFC3339,date)
	fmt.Println(t.Format(time.UnixDate))
}
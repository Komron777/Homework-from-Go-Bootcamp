// for testing the functions in package time (part of the learning)
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	unixTime := time.Unix(1000000000,45)
	fmt.Println(unixTime)
	///////////

	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	timeInUTC := time.Now().UTC()
	fmt.Println(timeInUTC.In(location))
}


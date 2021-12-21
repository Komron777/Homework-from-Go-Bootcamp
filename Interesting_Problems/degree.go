package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		h     int64
		m     int64
		d     int64
		which string
	)
	dt := time.Now()
	fmt.Printf("Current time: %02d:%02d\n", dt.Hour(), dt.Minute())
	fmt.Print("Do you want to rotate hour or minute arrow(h/m)? ")
	fmt.Scanln(&which)
	h, m = int64(dt.Hour()), int64(dt.Minute())
	switch which {
	case "h":
		fmt.Print("Degree: ")
		fmt.Scanln(&d)
		h += d / 30
		m += d % 30 * 2
	case "m":
		fmt.Print("Degree: ")
		fmt.Scanln(&d)
		h += d / 360
		m += d % 360 / 6
	}
	if m >= 60 {
		m -= 60
		h += 1
	}
	if h >= 24 {
		h -= 24
	}
	fmt.Printf("%02v:%02v\n", h, m)
}

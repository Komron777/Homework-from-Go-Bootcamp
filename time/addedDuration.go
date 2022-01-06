/*
https://stepik.org/lesson/359395/step/8?unit=343626

There is given date in format Unix-Time. 
We should input time in format "12 мин. 13 сек." which is duration.
This duration should be added to the given date, 
and then has to be converted to UnixDate format
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
	const now = 1589570165
	text,_ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Replace(text," мин. ","m",-1)
	text = strings.Replace(text," сек.","s",-1)
	text = strings.TrimSpace(text)
	t := time.Unix(now,0)
	dur,_ := time.ParseDuration(text)
	t = t.Add(dur)
	t = t.UTC()
	fmt.Println(t.Format(time.UnixDate))
}
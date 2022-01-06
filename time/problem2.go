/*
2020-05-15 08:00:00
If the time of the event is before lunchtime (13-00), then nothing needs to be changed, 
it is enough to display the date on the standard output in the same format.

If the event is to occur in the afternoon, you must reschedule it at the same time the next day, 
and then print it to standard output in the same format.


Sample Input:

2020-05-15 08:00:00

Sample Output:

2020-05-15 08:00:00


https://stepik.org/lesson/359395/step/4?unit=343626
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
	var date string
	
    date, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	date = strings.TrimSpace(date)
	t,_ := time.Parse("2006-01-02 15:04:05",date)

	if t.Hour() >= 13 {
		t = t.AddDate(0,0,1)
	} 
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
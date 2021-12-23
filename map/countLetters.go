package main
import (
	"fmt"
)
func countLetters(text string) map[string]int{
	countMap := map[string]int{}
	for _,v := range text {
		countMap[string(v)]++
	}
	return countMap
}
func main(){
	fmt.Println(countLetters("How many years should I spend on learning Go?"))
}
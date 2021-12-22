package main

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"

)

func RemoveDigitsOrder(text string) ([]string) {
	//var numbers [] int
	words := strings.Split(text," ")
	newOrder := []string{}
	
	for _,j := range words{
		for _,j1 := range j{
			if unicode.IsDigit(j1){
				index,err := strconv.Atoi(string(j1))	
				if err == nil{
					newWord := strings.Replace(j,string(j1),"",1)
					newOrder[index] = newWord
				}
				break
			}
		}
	}
	return newOrder
}
func main(){
	fmt.Println(RemoveDigitsOrder("Th0e 2is crazy4 wor1ld bec3oming"))

}
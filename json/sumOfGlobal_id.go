/*
Find the sum of all elements in global_id fields

Link to the JSON file: https://github.com/semyon-dev/stepik-go/tree/master/work_with_json
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)
func main() {
	var sum uint64
	filename := "stepik_info.json"
	var items []struct {
		Id uint64 `json:"global_id"`
	}
	f,_ := os.Open(filename)
	json.NewDecoder(f).Decode(&items)
	for _,item := range items {
		sum += item.Id
	}
	fmt.Println(sum)
}

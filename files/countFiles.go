
/*
This program is intented to walk through all FILES within a specifed root directory andm COUNT them
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)
var count int

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err 
	}

	if info.IsDir() {
		return nil 
	}
	count++
	return nil
}

func main() {
	const root = "." 

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}
}
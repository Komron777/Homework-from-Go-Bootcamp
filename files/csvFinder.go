package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"io/ioutil"
)

// func walkFunc(path string, info os.FileInfo,err error) {
// 	if info.
// }
func main() {
	const root = "."
	f,err := ioutil.ReadAll()
	err := filepath.Walk(root,walkFunc)
	
}
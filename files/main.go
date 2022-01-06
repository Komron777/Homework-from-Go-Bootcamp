/*for testing what I learned*/
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"bufio"
)

// 	"net/textproto"
// 	"strconv"
//

func main() {
	// for i := 1; i <= 3; i++ {
	// 	file, err := os.Create(strconv.Itoa(i) + "test.tx")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer file.Close()
	// }
	// os.Rename("2test.txt","4test.txt")
	// for i := 4; i >= 2; i-- {
	// 	os.Remove(strconv.Itoa(i) + "test.txt")
	// }

	////////////////////////////////////

	//text2 := "I am struggling with learning GO"
	// text,err := ioutil.ReadFile("smth.txt")
	// if err != nil {
	// 	panic("File may not exist!")
	// }
	// fmt.Println(string(text))

	/*Writing to the file, reading from it, and comparing them*/
	// dataForFile := []byte("Тестовая строка, предназначенная для записи в файл")

	// // Создаем новый файл и записываем в него данные dataForFile
	// if err := ioutil.WriteFile("test.txt", dataForFile, 0600); err != nil {
	// 	panic("ERROR")
	// }

	// // Читаем данные из того же файла
	// dataFromFile, err := ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	panic("ERROR")

	// }

	///////////////////////////

	/*Usage of ReadAll function*/
	// fmt.Printf("dataForFile == dataFromFile: %v\n", bytes.Equal(dataFromFile, dataForFile))
	// b := bytes.NewReader([]byte("Данные в объекте io.Reader"))
	// data, err := ioutil.ReadAll(b)
	// if err != nil {
	// 	// ...
	// }
	// fmt.Printf("%s\n", data) // Данные в объекте io.Reader

	////////////////////////////////////
	/*Calling files and directories from directories,
	and working woth them*/
	// filesFromDir,err := ioutil.ReadDir(".")
	// if err != nil{
	// 	panic("Directory path is not correct")
	// }
	// //fmt.Println(filesFromDir)
	// for _,files := range filesFromDir {
	// 	fmt.Println(files.Name())
	// 	fmt.Println(files.ModTime())
	// 	fmt.Println(files.IsDir())
	// 	fmt.Println(files.Mode())
	// 	fmt.Println(files.Size())
	// 	fmt.Println(files.Sys())
	// 	fmt.Println("///////////////////////")
	// }
	////////////////////////////////

	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	panic("PathError")
	// }
	// defer f.Close()
	// fmt.Println("check if defer is working")
	// file1, _ := os.Create("great.go")
	// file2, _ := os.Create("great.go")
	// //fmt.Println(file1.SetDeadline())
	// d := []byte("something")
	// file1.WriteString("sdlfjljgldfvlkdf")
	// fmt.Println(file1.WriteAt(d,0))
	// info1,_ := file1.Stat()
	// info2,_ := file2.Stat()
	// fmt.Println(os.SameFile(info1,info2))

	/*Bufio READER*/
	// 	file, err := os.Open("test.txt")
	// 	if err != nil {
	// 		panic("ERROR")
	// 	}
	// 	defer file.Close()
	// 	buf := make([]byte, 94)
	// 	rd := bufio.NewReader(file)
	// 	n, err := rd.Read(buf)
	// 	fmt.Println(n)

	// 	if err != nil && err != io.EOF {
	// 		panic("error")
	// 	}

	// fmt.Printf("прочитано %d байт: %s\n", n, buf) // прочитано 10 байт: bufio ...

	// s, err := rd.ReadString('\n') // читаем данные до разрыва абзаца ('\n')
	// fmt.Printf("%s\n", s)         // ... здесь будет строка

	/*Bufio WRITER*/
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	panic("Cannot create the file with such extension")
	// }
	// defer file.Close()
	// w := bufio.NewWriter(file)
	// n, err := w.WriteString("input string")
	// fmt.Printf("%d written\n", n)
	// w.Flush()

	/*Bufio SCANNER*/
	// var sum int
	// file, err := os.Open("numbers.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// s := bufio.NewScanner(file)
	// for s.Scan() {
	// 	num,_ := strconv.Atoi(s.Text())
	// 	sum += num
	// }
	// fmt.Println(sum)

	/*encoding/csv*/
	// Записывать данные, а в дальнейшем читать их мы будем из буфера,
	// но его можно заменить любым другим объектом, удовлетворяющим
	// интерфейсу io.ReadWriter
	// buf := bytes.NewBuffer(nil)

	// w := csv.NewWriter(buf)

	// for i := 1; i <= 3; i++ {
	// 	// Запись данных может производится поэтапно, например в цикле
	// 	val1 := fmt.Sprintf("row %d col 1", i)
	// 	val2 := fmt.Sprintf("row %d col 2", i)
	// 	val3 := fmt.Sprintf("row %d col 3", i)
	// 	if err := w.Write([]string{val1, val2, val3}); err != nil { // Аргументом Write является срез строк
	// 		// ...
	// 	}
	// }
	// w.Flush() // Этот метод приведет к фактической записи данных из буфера csv.Writer в buf

	// // Либо данные можно записать за один раз
	// w.WriteAll([][]string{ // Аргументом WriteAll является срез срезов строк
	// 	{"row 4 col 1", "row 4 col 2", "row 4 col 3"},
	// 	{"row 5 col 1", "row 5 col 2", "row 5 col 3"},
	// })

	// r := csv.NewReader(buf)

	// for i := 1; i <= 2; i++ {
	// 	// Читать данные мы тоже можем построчно, получая срез строк за каждую итерацию
	// 	_, err := r.Read()
	// 	if err != nil && err != io.EOF { // Здесь тоже нужно учитывать конец файла
	// 		// ...
	// 	}
	// 	//fmt.Println(row)
	// }

	// // Либо прочитать данные за один раз
	// data, err := r.ReadAll()
	// if err != nil {
	// 	// Когда мы читаем данные до конца файла io.EOF не возвращается, а служит сигналом к завершению чтения
	// 	// ...
	// }

	// for _, row := range data {
	// 	fmt.Println(row)
	// }

	// // [row 1 col 1 row 1 col 2 row 1 col 3]
	// // [row 2 col 1 row 2 col 2 row 2 col 3]
	// // [row 3 col 1 row 3 col 2 row 3 col 3]
	// // [row 4 col 1 row 4 col 2 row 4 col 3]
	// // [row 5 col 1 row 5 col 2 row 5 col 3]

	reader := bufio.NewReader(os.Stdin)
	text,_ := reader.ReadString('\n')
	fmt.Print(text)
}

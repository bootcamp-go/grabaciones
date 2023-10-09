package main

import (
	"fmt"
	"io"
	"os"
)

/*
	Notas del Orador:
	- Aplicando el manejo de errores tradicional de Go, el ejemplo anterior se veria asi...
*/
func main() {
	fileName := "test.txt"

	data, err := GetFileData(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// process data
	// ...
	data += " - processed"

	// print data
	fmt.Println(data)
}

// func GetFileData(fileName string) string {
// 	f, err := os.Open(fileName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	data, err := io.ReadAll(f)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return string(data)
// }


// refactoring
func GetFileData(fileName string) (data string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	dataBytes, err := io.ReadAll(f)
	if err != nil {
		return
	}

	data = string(dataBytes)
	return
}
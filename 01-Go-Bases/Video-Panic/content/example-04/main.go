package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "test.txt"

	data := GetFileData(fileName)

	// process data
	// ...
	data += " - processed"

	// print data
	fmt.Println(data)
}

func GetFileData(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(data)
}
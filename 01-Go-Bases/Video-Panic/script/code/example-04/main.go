package main

import (
	"fmt"
	"io"
	"os"
)

// Video-Speech: Aca tenemos un caso de un panic forzado donde el programa finaliza abruptamente
// al no poder abrir el archivo, el programa genera un panic
// si bien podemos usar el panic de esta forma, se puede aun aplicar el manejo de errores tradicional de Go
// hay casos como cuando se utilizan conexiones a bases de datos donde el panic tambien podría aplicarse si toda la app depende de la conexión
// pero en general, el manejo de errores es la forma correcta de hacerlo
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

// refactoring
func GetFileDataRefactored(fileName string) (data string, err error) {
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
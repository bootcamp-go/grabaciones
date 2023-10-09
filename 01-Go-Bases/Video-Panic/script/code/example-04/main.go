package main

import (
	"fmt"
	"io"
	"os"
)

/*
	Notas del Orador:
	- Asi como go genera panic para ciertos casos en forma nativa, como por ejemplo cuando se intenta acceder a un indice de un slice que no existe,
	nosotros tambien podemos definir reglas donde ciertos procesos o funciones de mayor nivel de nuestra applicación lo generen panic.
	- Hay casos donde se utilizan conexiones a bases de datos donde el panic tambien podría aplicarse si toda la app depende de la conexión.
	- Aclaración: si bien podemos usar el panic de esta forma, se puede y recomienda aun así aplicar el manejo de errores tradicional de Go.

	- Entrando en detalle en el ejemplo, en este caso generaremos panic manualmente en caso de no poder abrir un archivo.
*/
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
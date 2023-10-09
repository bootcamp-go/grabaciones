package main

import "fmt"

func main() {
	// number
	n1 := 1

	// get the square of a number of n1
	square := Square(n1)
	fmt.Printf("The square of %d is %d\n", n1, square)

	// get the square of a number of n2
	n2 := 2
	square = Square(n2)
	fmt.Printf("The square of %d is %d\n", n2, square)

	// get the square of a number of n3
	n3 := 3
	square = Square(n3)
	fmt.Printf("The square of %d is %d\n", n3, square)

	// get the square of the sum of n1 and n2
	square = Square(Sum(n1, n2))
	fmt.Printf("The square of the sum of %d and %d is %d\n", n1, n2, square)
}

/*
	Notas del Orador:
	- Algo que tambien podemos hacer es utilizar retornos nombrados, donde el parametro de retorno tiene un nombre,
	que representa una variable inicializada con el valor por defecto del type. En este caso el type es int, por lo que el valor por defecto es 0.
	- Luego durante el proceso, simplemente le asignamos el resultado a la variable result.
	- Finalmente, retornamos sin especificar, pues go ya sabe que queremos retornar la variable result. De todas formas si queremos podemos devolverlo concretamente.
*/
// Square returns the square of a number
func Square(n int) (result int) {
	result = n * n
	return
}

/*
	Notas del Orador:
	- Es similar a realizar lo siguiente
*/
func Square2(n int) int {
	var result int
	result = n * n
	return result
}

/*
	Notas del Orador:
	- Por ultimo creamos una función que se encargue de sumar dos números.
	- La entrada son dos números enteros, el proceso es sumarlos y la salida es un número entero, el resultado de la suma.
	- si tenemos multiples parametros de entrada o salida del mismo type, podemos especificar el type una sola vez al final.
*/
// Sum returns the sum of two numbers
func Sum(n1, n2 int) (result int) {
	result = n1 + n2
	return
}
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
}

/*
	Notas del Orador:
	- Para evitar repetir la misma lógica, podemos crear una función que se encargue de calcular el cuadrado de un número.
	- Primero indicamos el type func, luego el nombre de la función
	- Luego definimos las 3 partes de la función: la entrada, proceso y salida
	- Entrada: es un número entero, el proceso es multiplicar el número por si mismo, ponemos n y el type int
	- Proceso: es multiplicar el número por si mismo
	- Salida: es un número entero, el resultado de la multiplicación, ponemos int
*/
// Square returns the square of a number
func Square(n int) int {
	return n * n
}
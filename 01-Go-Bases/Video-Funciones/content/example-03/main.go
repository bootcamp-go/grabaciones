package main

import "fmt"

func main() {
	// number
	n := 10

	square := Square(n)
	fmt.Println("the square of", n, "is", square)

	n = 5
	square = Square(n)
	fmt.Println("the square of", n, "is", square)

	n = 2
	square = Square(n)
	fmt.Println("the square of", n, "is", square)

	// calculate the square of the sum of two numbers
	result := Square(Sum(2, 3))
	fmt.Println("the square of the sum of 2 and 3 is", result)
}

func Square(n int) (result int) {
	result = n * n
	return
}

func Sum(n1, n2 int) (result int) {
	result = n1 + n2
	return
}
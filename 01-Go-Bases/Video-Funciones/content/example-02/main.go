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
}

func Square(n int) (result int) {
	result = n * n
	return
}
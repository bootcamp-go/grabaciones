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

// Square returns the square of a number
func Square(n int) (result int) {
	result = n * n
	return
}

// Sum returns the sum of two numbers
func Sum(n1, n2 int) (result int) {
	result = n1 + n2
	return
}
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

// Square returns the square of a number
func Square(n int) int {
	return n * n
}
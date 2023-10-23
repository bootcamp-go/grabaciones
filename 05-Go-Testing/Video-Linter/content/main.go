package main

import (
	"errors"
	"math"
)

func main() {
	// Create a new Calculator.
	c := &Calculator{}

	// Add 2 and 3.
	result := c.Add(2, 3)
	println(result)
}

// calculator.go
var (
	// DivisionError is returned when a division by zero is attempted.
	DivisionError = errors.New("division by zero")
	// ErrCalculatorSquareRoot is returned when a square root of a negative number is attempted.
	ErrCalculatorSquareRoot = errors.New("Square root of negative number.")
)

// Calculator provides methods to perform simple calculations.
type Calculator struct{}

// Add returns the sum of two integers.
func (c *Calculator) Add(a, b float64) (result float64) {
	result = a + b
	return
}

// Subtract returns the difference of two integers.
func (c *Calculator) Subtract(a, b float64) (result float64) {
	result = a - b
	return
}

// Divide returns the quotient of two integers.
func (c *Calculator) Divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = DivisionError
		return
	}
	result = a / b
	return
}

// SquareRoot returns the square root of an integer.
func (c *Calculator) SquareRoot(a float64) (result float64, err error) {
	if a < 0 {
		err = ErrCalculatorSquareRoot
		return
	}
	result = math.Sqrt(a)
	return
}
package main

import "fmt"

func main() {
	One()
}

func One() {
	fmt.Println("I'm function one")
	Two()
}

func Two() {
	fmt.Println("I'm function two")
	panic("Panic in function two")
}
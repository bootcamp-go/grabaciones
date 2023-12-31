package main

import (
	"fmt"
	"time"
)

// SauceMaker makes the sauce
// It takes 13 seconds to make the sauce
func SauceMaker() string {
	startTime := time.Now()
	fmt.Println("Making the sauce: heating olive oil ")
	time.Sleep(1 * time.Second)
	fmt.Println("Making the sauce: browing the garlic ") 
	time.Sleep(2 * time.Second)
	fmt.Println("Making the sauce: browing the onions")
	time.Sleep(2 * time.Second)
	fmt.Println("Making the sauce: cooking the meatballs")
	time.Sleep(5 * time.Second)
	fmt.Println("Making the sauce: cooking the diced tomatoes")
	time.Sleep(3 * time.Second)
	endTime := time.Since(startTime).Seconds()
	
	fmt.Printf("making the sauce took %f seconds\n", endTime)

	return "sauce"
}

// PastaMaker makes the pasta
// It takes 12 seconds to make the pasta
func PastaMaker() string {
	startTime := time.Now()
	fmt.Println("Making the pasta: boiling water")
	time.Sleep(8 * time.Second)
	fmt.Println("Making the pasta: cooking pasta")
	time.Sleep(4 * time.Second)
	endTime := time.Since(startTime).Seconds()

	fmt.Printf("making the pasta took %f seconds\n", endTime)

	return "cooked pasta"
}

func main() {
	startTime := time.Now()
	sauce := SauceMaker()
	pasta := PastaMaker()
	endTime := time.Since(startTime).Seconds()

	fmt.Printf("%s and %s are done\n", sauce, pasta)
	fmt.Printf("cooking dinner took %f seconds\n", endTime)
}

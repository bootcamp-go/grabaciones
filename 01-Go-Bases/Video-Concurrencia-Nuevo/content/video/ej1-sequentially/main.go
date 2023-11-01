package main

import (
	"fmt"
	"time"
)

// SauceMaker makes the sauce
// It takes 13 seconds to make the sauce
func SauceMaker() string {
	startTime := time.Now().UnixMilli()
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
	endTime := time.Now().UnixMilli()
	fmt.Printf("making the sauce took %f seconds\n", (float64(endTime)-float64(startTime))/1000)

	return "sauce"
}

// PastaMaker makes the pasta
// It takes 12 seconds to make the pasta
func PastaMaker() string {
	startTime := time.Now().UnixMilli()
	fmt.Println("Making the pasta: boiling water")
	time.Sleep(8 * time.Second)
	fmt.Println("Making the pasta: cooking pasta")
	time.Sleep(4 * time.Second)
	endTime := time.Now().UnixMilli()
	fmt.Printf("making the pasta took %f seconds\n", (float64(endTime)-float64(startTime))/1000)

	return "cooked pasta"
}

func main() {
	startTime := time.Now().UnixMilli()
	sauce := SauceMaker()
	pasta := PastaMaker()
	endTime := time.Now().UnixMilli()

	fmt.Printf("%s and %s are done\n", sauce, pasta)
	fmt.Printf("cooking dinner took %f seconds\n", (float64(endTime)-float64(startTime))/1000)
}

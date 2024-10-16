package main

import (
	"fmt"
)

func main() {
	// If statement
	age := 18
	if age < 18 {
		fmt.Println("You are a minor")
	} else {
		fmt.Println("You are an adult")
	}

	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}
}

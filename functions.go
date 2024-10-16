package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(5, 3)
	fmt.Printf("The sum of 5 and 3 is %d\n", sum)
}

package main

import (
	"fmt"
)

func main() {
	total := addNumbers(5, 10)
	fmt.Println("Total:", total)
}

func addNumbers(a int, b int) int {
	sum := a + b
	fmt.Println("Sum is:", sum)
	return sum
}

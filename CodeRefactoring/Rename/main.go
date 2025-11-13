package main

import (
	"Mi"
	"fmt"
)

func main() {
	r := 5.0
	area := palit(r) // works because same package
	fmt.Println("Area:", area)
}

func addNumbers(a int, b int) int {
	sum := a + b
	fmt.Println("Sum is:", sum)
	return sum
}

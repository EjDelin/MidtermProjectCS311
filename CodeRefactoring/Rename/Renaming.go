package main

import (
	"MidProjectPresentation/CodeRefactoring/CodeInspection"
	"fmt"
)

func main() {
	r := 5.0
	area := CodeInspection.TestR(r) // works because same package
	fmt.Println("Area:", area)
}

// AddNumbers TODO: Call this method in other classes
func AddNumbers(a int, b int) int {
	sum := a + b
	fmt.Println("Sum is:", sum)
	return sum
}

package main

import "fmt"

func main() {
	r := 5.0
	area := testR(r) // works because same package
	fmt.Println("Area:", area)
}

func addNumbers(a int, b int) int {
	sum := a + b
	fmt.Println("Sum is:", sum)
	return sum
}

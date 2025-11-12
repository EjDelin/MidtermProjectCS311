package main

import "fmt"

func main() {
	total := calculateSum(5, 10, 0, 14)
	displayTotal(total)
}

func displayTotal(total int) (int, error) {
	return fmt.Println("Total:", total)
}

func calculateSum(a int, b int, c int, d int) int {
	sum := a + b
	fmt.Println("Sum is:", sum)
	return sum
}

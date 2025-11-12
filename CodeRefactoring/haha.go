package main

import "fmt"

func main() {
	total := calculateSum(5, 10, 0, 14)
	//unusedValue := 123 // GoLand will highlight this line
	fmt.Println("Total:", total)
}

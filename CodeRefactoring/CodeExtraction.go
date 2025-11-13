package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 10.0
	area := math.Pi * math.Pow(radius, 2)
	perimeter := 2 * math.Pi * radius

	fmt.Printf("Circle Area: %.2f\n", area)
	fmt.Printf("Circle Perimeter: %.2f\n", perimeter)
}

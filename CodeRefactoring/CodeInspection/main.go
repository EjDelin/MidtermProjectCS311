package main

import (
	"fmt"
	"math"
)

func calculateArea(radius float64) float64 {
	if radius <= 0 {
		fmt.Println("Invalid radius")
		return 0
	}
	return math.Pi * radius * radius
}

func calculateDiameter(radius float64) float64 {
	return radius * 2
}

func main() {
	r := -5.0
	area := calculateArea(r)
	diameter := calculateDiameter(r)
	//unusedValue := 100 // Warning: unused variable
	fmt.Printf("Area: %.2f, Diameter: %.2f\n", area, diameter)
}

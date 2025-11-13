package main

import "fmt"

//go:generate go run scripts/gen.go
//go:generate go run scripts/vehicle.go

func main() {
	fmt.Println(GetMotivationalQuote())

}

func add(s int, b int) int {
	return s + b

}

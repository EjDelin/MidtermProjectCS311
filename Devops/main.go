package main

//go:generate go run scripts/gen.go

import "fmt"

func main() {
	// Call the function that will be auto-generated
	fmt.Println(GetMotivationalQuote())
}

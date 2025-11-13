package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	//Basic Completion
	message := greet("Renzo")
	fmt.Println(message)

	//Type-aware
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)

}

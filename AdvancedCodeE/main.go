package main

import (
	"fmt"
	"strings"
	"time"
)

type Car struct {
	Model string
	Year  int
	Brand string
}

// Person represents a simple person.
type Person struct {
	FirstName string
	LastName  string
	BirthYear int
}

// FullName returns the full name.
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// Age calculates the person's age.
func (p Person) Age() int {
	// 2. DEMO (IntelliSense - Module):
	// Basic Code Completion with time module
	age := time.Now().Year() - p.BirthYear
	return age
}

// NAVIGATION

// Greet prints a greeting.
func (p Person) Greet() {
	// 3. DEMO (Navigation - Go to Definition):
	fmt.Printf("Hello, my name is %s and I am %d years old!\n",
		p.FullName(), p.Age())
}

// Greeter is an interface for things that can greet.
type Greeter interface {
	// 4. DEMO (Navigation - Find Usages):
	Greet()
}

// SayHello makes a Greeter greet.
func SayHello(g Greeter) {
	g.Greet()

}

//Generate - Implement Methods
//Generate - type From JSON

func main() {
	var name = "GoLand"
	fmt.Println("Welcome to", name)

	// 5. DEMO (IntelliSense - Struct Literal):
	// IntelliSense
	user := Person{
		FirstName: "James",
		LastName:  "Bond",
		BirthYear: 2000,
	}

	// 6. DEMO (Code Completion):
	user.Greet()
	user.Age()
	user.FullName()

	// 7. DEMO (IntelliSense - Module):
	// strings
	text := "goland IDE demo"
	lower := strings.ToLower(text)
	fmt.Println(lower)
	// 8. DEMO (Shortcuts - Commenting):
	// Select these next two lines of code with your mouse.
	// Press 'Ctrl+/' (or 'Cmd+/') to toggle them as comments.
	fmt.Println("This is a test line 1.")
	fmt.Println("This is a test line 2.")

	// 9. DEMO (Shortcuts - Formatting):
	//unformatted comment
	fmt.Println("This line will be auto-formatted.")

}

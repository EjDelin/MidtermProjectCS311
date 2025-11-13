package main

import (
	"fmt"
	"math" // 1. For IntelliSense on a modul
)

// User represents a simple user.
type User struct {
	// 2. DEMO: Start typing 'Username' here and see completion
	Username string
	Email    string
	IsActive bool
}

// Greet prints a greeting message.
func (u *User) Greet() {
	// 3. DEMO: Type 'u.' to see 'Username', 'Email', etc.
	fmt.Printf("Hello, %s! Welcome back.\n", u.Username)
}

// GetProfileInfo gets a formatted string of user info.
func (u *User) GetProfileInfo() string {
	userInfo := fmt.Sprintf("User: %s", u.Username)

	// 4. DEMO (Navigation): Ctrl+Click (or F12) on 'formatUserEmail'
	//    to "Go to Definition"
	emailInfo := formatUserEmail(u.Email)

	return fmt.Sprintf("%s\n%s", userInfo, emailInfo)
}

// formatUserEmail formats the email string for display.
func formatUserEmail(emailString string) string {
	// 5. DEMO (Navigation): Right-click this function name -> "Find Usages"
	//    (or Alt+F7)
	return fmt.Sprintf("Email: <%s>", emailString)
}

// NewUser is a constructor-like function for User.
func NewUser(username, email string) *User {
	// 6. DEMO (IntelliSense):
	//    When filling this struct, GoLand suggests the field names
	return &User{
		Username: username,
		Email:    email,
		IsActive: true,
	}
}

func main() {
	// 7. DEMO (IntelliSense):
	//    Start typing 'NewUser(' and see the parameter hints (username, email)
	user1 := NewUser("go_developer", "gopher@example.com")

	// 8. DEMO (Code Completion):
	//    Type 'user1.' and see the list of methods and fields
	//    (Greet, GetProfileInfo, Username, etc.)
	user1.Greet()

	profile := user1.GetProfileInfo()
	fmt.Println(profile)

	// 9. DEMO (Shortcuts - Commenting):
	//    Select the two lines below and press Ctrl+/ (or Cmd+/)
	//    to toggle line comments.
	fmt.Println("This is a test line 1.")
	fmt.Println("This is a test line 2.")

	// 10. DEMO (Shortcuts - Formatting):
	//     Delete the indentation for 'val := 1+2' to make it messy
	//     Then, press the "Reformat Code" shortcut (Ctrl+Alt+L or Cmd+Option+L)
	val := 1 + 2
	fmt.Printf("The value is: %d\n", val)

	// 11. DEMO (IntelliSense - Imports):
	//     Type 'math.' and see all the functions from the imported module
	fmt.Printf("The value of Pi is: %f\n", math.Pi)
}

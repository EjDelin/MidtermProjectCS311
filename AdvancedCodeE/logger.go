package main

import "fmt"
import _ "errors"

// Logger interface defines logging operations.
type Logger interface {
	Log(message string)
	Error(err error)
}

// MyLogger is a concrete implementation of the Logger.
type MyLogger struct {
}

// Log implements the Logger interface.
func (m *MyLogger) Log(message string) {
	// Implement a simple log to console
	fmt.Printf("[LOG] %s\n", message)
}

// Error implements the Logger interface.
func (m *MyLogger) Error(err error) {
	// Implement a simple error log to console
	fmt.Printf("[ERROR] %s\n", err.Error())
}

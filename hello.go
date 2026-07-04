package main

import "fmt"

// Simple greeting function
func greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go.", name)
}

func main() {
	fmt.Println(greet("World"))
}

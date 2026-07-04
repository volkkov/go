package main

import "strings"

// IsValidEmail does a basic check for @ and . in the string
func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

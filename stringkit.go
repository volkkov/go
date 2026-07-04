package main

import (
	"strings"
	"unicode"
)

// Capitalize makes the first letter of a string uppercase
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// SlugifyLite converts a string into a simple URL-friendly slug
func SlugifyLite(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	var b strings.Builder
	lastDash := false
	for _, r := range s {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			lastDash = false
		case r == ' ' || r == '-' || r == '_':
			if !lastDash {
				b.WriteRune('-')
				lastDash = true
			}
		}
	}
	return strings.Trim(b.String(), "-")
}

// CountWords returns the number of words in a string
func CountWords(s string) int {
	fields := strings.Fields(s)
	return len(fields)
}

// TruncateWithEllipsis shortens a string to maxLen characters, adding "..." if truncated
func TruncateWithEllipsis(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}

// IsBlank checks whether a string is empty or contains only whitespace
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

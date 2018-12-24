// Package twofer implements the "Two Fer" exercise.
package twofer

import "fmt"

// ShareWith returns the string "One for X, one for me." where X depends on the input.
func ShareWith(name string) string {
	var other string
	if name == "" {
		other = "you"
	} else {
		other = name
	}
	return fmt.Sprintf("One for %s, one for me.", other)
}

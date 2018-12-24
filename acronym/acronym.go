// Package acronym implements the "acronym" exercise.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns the acronym of the given string
func Abbreviate(s string) string {
	words := regexp.MustCompile(`[\s-]+`).Split(strings.ToUpper(s), -1)
	firstLetters := make([]string, len(words))
	for i, v := range words {
		firstLetters[i] = string(v[0])
	}

	return strings.Join(firstLetters, "")
}

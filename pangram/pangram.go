// Package pangram implements the "Pangram" exercise.
package pangram

import "unicode"
import "strings"

// IsPangram checks whether the given string is a pangram
func IsPangram(in string) bool {
	in = strings.ToLower(in)
	seen := make(map[rune]bool, 26)
	for _, r := range in {
		if !unicode.IsLetter(r) {
			continue
		}
		seen[r] = true
	}
	return len(seen) == 26
}

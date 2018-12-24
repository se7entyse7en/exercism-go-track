// Package isogram implements the "Isogram" exercise.
package isogram

import "unicode"

// IsIsogram checks whether the given input is isogram.
func IsIsogram(s string) bool {
	m := make(map[rune]bool, len(s))
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToUpper(r)
		if m[r] {
			return false
		}
		m[r] = true
	}
	return true
}

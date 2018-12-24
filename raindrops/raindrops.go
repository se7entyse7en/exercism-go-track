// Package raindrops implements the "Raindrops" exercise.
package raindrops

import (
	"strconv"
)

// Convert returns a string that is dependant on the factors of the given int number
func Convert(n int) (result string) {
	if n%3 == 0 {
		result += "Pling"
	}
	if n%5 == 0 {
		result += "Plang"
	}
	if n%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(n)
	}

	return
}

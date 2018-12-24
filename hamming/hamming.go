// Package hamming implements the "Hamming" exercise.
package hamming

import "errors"

// Distance returns the Hamming distance between two strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings `a` and `b` has different sizes")
	}

	counter := 0
	for i := range a {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}

// Package reverse implements the "Reverse String" exercise.
package reverse

// String returns the input string inverted.
func String(in string) string {
	runes := []rune(in)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

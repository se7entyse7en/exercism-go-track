// Package accumulate implements the "Accumulate" exercise.
package accumulate

// Accumulate returns the given collection whose elements are transformed with the given function
func Accumulate(collection []string, converter func(string) string) []string {
	transformed := make([]string, len(collection))
	for i, v := range collection {
		transformed[i] = converter(v)
	}
	return transformed
}

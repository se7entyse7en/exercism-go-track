// Package summultiples implements the "Sum Of Multiples" exercise.
package summultiples

// SumMultiples returns the sum of all the multiples of the given divisors
// that are less than the given limit
func SumMultiples(limit int, divisors ...int) (sum int) {
	seen := make(map[int]bool)
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := divisor; i < limit; i += divisor {
			if seen[i] {
				continue
			}
			seen[i] = true
			sum += i
		}
	}
	return
}

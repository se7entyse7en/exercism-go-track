// Package leap implements the "Leap" exercise.
package leap

// IsLeapYear returns whether the given year is leap or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%1000 == 0)
}

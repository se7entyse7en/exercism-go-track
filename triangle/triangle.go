// Package triangle implements the "Triangle" exercise.
package triangle

import "math"

// Kind indicates the triangle type
type Kind int

const (
	// NaT represents a not a triangle
	NaT Kind = iota
	// Equ represents an equilateral triangle
	Equ
	// Iso represents an isosceles triangle
	Iso
	// Sca represents an scalene triangle
	Sca
)

// KindFromSides returns the kind of a triangle given its length
func KindFromSides(a, b, c float64) (k Kind) {
	validSides := a > 0 && b > 0 && c > 0 &&
		!math.IsInf(a, 1) && !math.IsInf(b, 1) && !math.IsInf(c, 1)
	triangleInequality := (a+b >= c) && (a+c >= b) && (b+c >= a)
	isTriangle := validSides && triangleInequality
	if !isTriangle {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a != b && b != c && a != c {
		k = Sca
	} else {
		k = Iso
	}
	return
}

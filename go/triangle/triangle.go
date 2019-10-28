// Package triangle determines if a triangle is equilateral, isosceles, or scalene.
package triangle

import "math"

// Kind data type is a string
type Kind int

const (
	// NaT is not a triangle
	NaT = iota
	// Equ is an equilateral triangle
	Equ
	// Iso is an isosceles triangle
	Iso
	// Sca is a scalene triangle
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Checks to make sure each side is greater than 0 and not equal to positive infinity
	if a > 0 && a != math.Inf(1) && b > 0 && b != math.Inf(1) && c > 0 && c != math.Inf(1) {
		// The sum of the lengths of two sides must be greater or equal to the length of the third side.
		if a+b >= c && a+c >= b && b+c >= a {
			// If all sides are equal
			if a == b && b == c {
				return Equ
			}
			// If two sides are equal
			if a == b || a == c || b == c {
				return Iso
			}

			// If no sides are equal
			if a != b && b != c {
				return Sca
			}
		}
	}
	return NaT
}

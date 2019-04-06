// Package fuelconv is a sample package.
package fuelconv

type (
	Liter  float64
	Gallon float64
)

// LToG converts a Liter into a Gallon.
func (l Liter) LToG() Gallon {
	return Gallon(l / 3.785411784)
}

// GToL converts a Gallon into a Liter.
func (g Gallon) GToL() Liter {
	return Liter(g * 3.785411784)
}

package triangle

import "math"

const testVersion = 3

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	switch {
	case
		math.IsNaN(a), math.IsNaN(b), math.IsNaN(c),
		math.IsInf(a, 0), math.IsInf(b, 0), math.IsInf(c, 0),
		a == 0 && b == 0 && c == 0,
		a < 0, b < 0, c < 0,
		a+b < c, a+c < b, b+c < a:
		return NaT
	case a == b && a == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}

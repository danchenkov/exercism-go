package space

import (
	"math"
)

type Planet string

var orbitalPeriods = map[Planet]float64{
	"Earth":   31557600,
	"Mercury": 0.2408467 * 31557600,
	"Venus":   0.61519726 * 31557600,
	"Mars":    1.8808158 * 31557600,
	"Jupiter": 11.862615 * 31557600,
	"Saturn":  29.447498 * 31557600,
	"Uranus":  84.016846 * 31557600,
	"Neptune": 164.79132 * 31557600,
}

func round(value float64, precision int) float64 {
	var r float64
	pow := math.Pow(10, float64(precision))
	digit := pow * value
	_, div := math.Modf(digit)
	if div >= 0.5 {
		r = math.Ceil(digit)
	} else {
		r = math.Floor(digit)
	}
	return r / pow
}

func Age(seconds float64, planet Planet) float64 {
	return round(seconds/orbitalPeriods[planet], 2)
}

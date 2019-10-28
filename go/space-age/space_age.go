// Package space determines your age in years on a specific planet
package space

import (
	"math"
)

type Planet string
const earthSeconds = float64(31557600)
var planets = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) (age float64) {
	tempAge := (seconds / earthSeconds) / planets[planet]
	age = math.Round(tempAge * 100) / 100
	return age
}
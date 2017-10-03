package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

type Kind string

var NaT Kind = "NaT" // not a triangle
var Equ Kind = "Equ" // equilateral
var Iso Kind = "Iso" // isosceles
var Sca Kind = "Sca" // scalene

func validNumber(floats []float64) bool {
	for _, val := range floats {
		if val <= 0 || math.IsNaN(val) || math.IsInf(val, 0) {
			return false
		}
	}
	return true
}

func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	if sides[0]+sides[1] < sides[2] || !validNumber(sides) {
		return NaT
	}
	if sides[0] == sides[1] && sides[1] == sides[2] {
		return Equ
	}
	if sides[0] == sides[1] || sides[1] == sides[2] || sides[0] == sides[2] {
		return Iso
	}
	return Sca
}

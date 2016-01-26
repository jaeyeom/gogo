// Package pattern is an example for Go functions. The concept of
// coding is from Section 1.3 of "Structure and Interpretation of
// Computer Programs, second edition, by Harold Abelson, Gerald
// Sussman and Julie Sussman", published by The MIT Press.
package pattern

import "math"

type Func func(float64) float64
type Transform func(Func) Func

const tolerance = 0.00001
const dx = 0.00001

func Square(x float64) float64 {
	return x * x
}

func FixedPoint(f Func, firstGuess float64) float64 {
	closeEnough := func(v1, v2 float64) bool {
		return math.Abs(v1-v2) < tolerance
	}
	var try Func
	try = func(guess float64) float64 {
		next := f(guess)
		if closeEnough(guess, next) {
			return next
		} else {
			return try(next)
		}
	}
	return try(firstGuess)
}

func FixedPointOfTransform(g Func, transform Transform, guess float64) float64 {
	return FixedPoint(transform(g), guess)
}

func Deriv(g Func) Func {
	return func(x float64) float64 {
		return (g(x+dx) - g(x)) / dx
	}
}

func NewtonTransform(g Func) Func {
	return func(x float64) float64 {
		return x - (g(x) / Deriv(g)(x))
	}
}

func Sqrt(x float64) float64 {
	return FixedPointOfTransform(func(y float64) float64 {
		return Square(y) - x
	}, NewtonTransform, 1.0)
}

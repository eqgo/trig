package trig

import "math"

// Sin returns the sine of the radian argument x.
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Cos returns the cosine of the radian argument x.
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Tan returns the tangent of the radian argument x.
func Tan(x float64) float64 {
	return math.Tan(x)
}

// Sec returns the secant of the radian argument x.
func Sec(x float64) float64 {
	return 1 / Cos(x)
}

// Csc returns the cosecant of the radian argument x.
func Csc(x float64) float64 {
	return 1 / Sin(x)
}

// Cot returns the cotangent of the radian argument x.
func Cot(x float64) float64 {
	return 1 / Tan(x)
}

// Arcsin returns the inverse sine, in radians, of x.
func Arcsin(x float64) float64 {
	return math.Asin(x)
}

// Arccos returns the inverse cosine, in radians, of x.
func Arccos(x float64) float64 {
	return math.Acos(x)
}

// Arctan returns the inverse tangent, in radians, of x.
func Arctan(x float64) float64 {
	return math.Atan(x)
}

// Arcsec returns the inverse secant, in radians, of x.
func Arcsec(x float64) float64 {
	return Arccos(1 / x)
}

// Arccsc returns the inverse cosecant, in radians, of x.
func Arccsc(x float64) float64 {
	return Arcsin(1 / x)
}

// Arccot returns the inverse cotangent, in radians, of x.
func Arccot(x float64) float64 {
	y := Arctan(1 / x)
	if x < 0 {
		y += Pi
	}
	return y
}

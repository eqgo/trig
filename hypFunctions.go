package trig

import "math"

// Sinh returns the hyperbolic sine of x.
func Sinh(x float64) float64 {
	return math.Sinh(x)
}

// Cosh returns the hyperbolic cosine of x.
func Cosh(x float64) float64 {
	return math.Cosh(x)
}

// Tanh returns the hyperbolic tangent of x.
func Tanh(x float64) float64 {
	return math.Tanh(x)
}

// Sech returns the hyperbolic secant of x.
func Sech(x float64) float64 {
	return 1 / Cosh(x)
}

// Csch returns the hyperbolic cosecant of x.
func Csch(x float64) float64 {
	return 1 / Sinh(x)
}

// Coth returns the hyperbolic cotangent of x.
func Coth(x float64) float64 {
	return 1 / Tanh(x)
}

// Arcsinh returns the inverse hyperbolic sine of x.
func Arcsinh(x float64) float64 {
	return math.Asinh(x)
}

// Arccosh returns the inverse hyperbolic cosine of x.
func Arccosh(x float64) float64 {
	return math.Acosh(x)
}

// Arctanh returns the inverse hyperbolic tangent of x.
func Arctanh(x float64) float64 {
	return math.Atanh(x)
}

// Arcsech returns the inverse hyperbolic secant of x.
func Arcsech(x float64) float64 {
	return Arccosh(1 / x)
}

// Arccsch returns the inverse hyperbolic cosecant of x.
func Arccsch(x float64) float64 {
	return Arcsinh(1 / x)
}

// Arccoth returns the inverse hyperbolic cotangent of x.
func Arccoth(x float64) float64 {
	return Arctanh(1 / x)
}

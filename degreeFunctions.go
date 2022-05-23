package trig

// SinDegrees returns the sine of the degree argument x.
func SinDegrees(x float64) float64 {
	return Sin(Radians(x))
}

// CosDegrees returns the cosine of the degree argument x.
func CosDegrees(x float64) float64 {
	return Cos(Radians(x))
}

// TanDegrees returns the tangent of the degree argument x.
func TanDegrees(x float64) float64 {
	return Tan(Radians(x))
}

// SecDegrees returns the secant of the degree argument x.
func SecDegrees(x float64) float64 {
	return 1 / Cos(Radians(x))
}

// CscDegrees returns the cosecant of the degree argument x.
func CscDegrees(x float64) float64 {
	return 1 / Sin(Radians(x))
}

// CotDegrees returns the cotangent of the degree argument x.
func CotDegrees(x float64) float64 {
	return 1 / Tan(Radians(x))
}

// ArcsinDegrees returns the inverse sine, in degrees, of x.
func ArcsinDegrees(x float64) float64 {
	return Degrees(Arcsin(x))
}

// ArccosDegrees returns the inverse cosine, in degrees, of x.
func ArccosDegrees(x float64) float64 {
	return Degrees(Arccos(x))
}

// ArctanDegrees returns the inverse tangent, in degrees, of x.
func ArctanDegrees(x float64) float64 {
	return Degrees(Arctan(x))
}

// ArcsecDegrees returns the inverse secant, in degrees, of x.
func ArcsecDegrees(x float64) float64 {
	return Degrees(Arccos(1 / x))
}

// ArccscDegrees returns the inverse cosecant, in degrees, of x.
func ArccscDegrees(x float64) float64 {
	return Degrees(Arcsin(1 / x))
}

// ArccotDegrees returns the inverse cotangent, in degrees, of x.
func ArccotDegrees(x float64) float64 {
	y := Degrees(Arctan(1 / x))
	if x < 0 {
		y += 180
	}
	return y
}

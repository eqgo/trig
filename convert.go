package trig

// Pi is the value of the constant pi
const Pi = 3.14159265358979323846264338327950288419716939937510582097494459

// Radians returns the radian value of the inputted degrees.
func Radians(degrees float64) float64 {
	return degrees * Pi / 180
}

// Degrees returns the degree value of the inputted radians.
func Degrees(radians float64) float64 {
	return radians * 180 / Pi
}

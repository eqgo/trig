package trig

import "testing"

type basicTest struct {
	input  float64
	output float64
}

func TestRadians(t *testing.T) {
	tests := []basicTest{{180, Pi}, {45, Pi / 4}, {540, 3 * Pi}}
	testFunc(Radians, tests, "Radians", t)
}

func TestDegrees(t *testing.T) {
	tests := []basicTest{{Pi, 180}, {Pi / 4, 45}, {3 * Pi, 540}}
	testFunc(Degrees, tests, "Degrees", t)
}

func testFunc(f func(float64) float64, tests []basicTest, name string, t *testing.T) {
	for _, d := range tests {
		if f(d.input) != d.output {
			t.Errorf("%v(%v) should be %v, not %v", name, d.input, d.output, f(d.input))
		}
	}
}

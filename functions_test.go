package trig

import (
	"testing"
)

func TestFunctions(t *testing.T) {
	tests := []basicTest{{0, 0}, {Pi / 2, 1}}
	testFunc(Sin, tests, "Sin", t)

	tests = []basicTest{{0, 1}, {Pi, -1}}
	testFunc(Cos, tests, "Cos", t)

	tests = []basicTest{{0, 0}, {Pi / 4, 1}}
	testFunc(Tan, tests, "Tan", t)

	tests = []basicTest{{0, 1}, {Pi, -1}}
	testFunc(Sec, tests, "Sec", t)

	tests = []basicTest{{Pi / 2, 1}, {Pi / 6, 2}}
	testFunc(Csc, tests, "Csc", t)

	tests = []basicTest{{Pi / 4, 1}, {-Pi / 4, -1}}
	testFunc(Cot, tests, "Cot", t)

	tests = []basicTest{{1, Pi / 2}, {0, 0}}
	testFunc(Arcsin, tests, "Arcsin", t)

	tests = []basicTest{{1, 0}, {0, Pi / 2}}
	testFunc(Arccos, tests, "Arccos", t)

	tests = []basicTest{{1, Pi / 4}, {0, 0}}
	testFunc(Arctan, tests, "Arctan", t)

	tests = []basicTest{{1, 0}, {-1, Pi}}
	testFunc(Arcsec, tests, "Arcsec", t)

	tests = []basicTest{{1, Pi / 2}, {-1, -Pi / 2}}
	testFunc(Arccsc, tests, "Arccsc", t)

	tests = []basicTest{{-1, 3 * Pi / 4}, {1, Pi / 4}}
	testFunc(Arccot, tests, "Arccot", t)
}

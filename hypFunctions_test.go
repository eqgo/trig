package trig

import (
	"math"
	"testing"
)

func TestHypFunctions(t *testing.T) {
	tests := []basicTest{{0, 0}, {math.Inf(-1), math.Inf(-1)}}
	testFunc(Sinh, tests, "Sinh", t)

	tests = []basicTest{{0, 1}, {math.Inf(-1), math.Inf(1)}}
	testFunc(Cosh, tests, "Cosh", t)

	tests = []basicTest{{0, 0}, {math.Inf(-1), -1}}
	testFunc(Tanh, tests, "Tanh", t)

	tests = []basicTest{{0, 1}, {math.Inf(-1), 0}}
	testFunc(Sech, tests, "Sech", t)

	tests = []basicTest{{math.Inf(1), 0}, {math.Inf(-1), 0}}
	testFunc(Csch, tests, "Csch", t)

	tests = []basicTest{{math.Inf(1), 1}, {math.Inf(-1), -1}}
	testFunc(Coth, tests, "Coth", t)

	tests = []basicTest{{0, 0}, {math.Inf(-1), math.Inf(-1)}}
	testFunc(Arcsinh, tests, "Arcsinh", t)

	tests = []basicTest{{1, 0}, {math.Inf(1), math.Inf(1)}}
	testFunc(Arccosh, tests, "Arccosh", t)

	tests = []basicTest{{0, 0}, {1, math.Inf(1)}}
	testFunc(Arctanh, tests, "Arctanh", t)

	tests = []basicTest{{1, 0}, {0, math.Inf(1)}}
	testFunc(Arcsech, tests, "Arcsech", t)

	tests = []basicTest{{math.Inf(1), 0}, {math.Inf(-1), 0}}
	testFunc(Arccsch, tests, "Arccsch", t)

	tests = []basicTest{{1, math.Inf(1)}, {-1, math.Inf(-1)}}
	testFunc(Arccoth, tests, "Arccoth", t)
}

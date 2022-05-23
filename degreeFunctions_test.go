package trig

import (
	"testing"
)

func TestDegreeFunctions(t *testing.T) {
	tests := []basicTest{{0, 0}, {90, 1}}
	testFunc(SinDegrees, tests, "SinDegrees", t)

	tests = []basicTest{{0, 1}, {180, -1}}
	testFunc(CosDegrees, tests, "CosDegrees", t)

	tests = []basicTest{{0, 0}, {45, 1}}
	testFunc(TanDegrees, tests, "TanDegrees", t)

	tests = []basicTest{{0, 1}, {180, -1}}
	testFunc(SecDegrees, tests, "SecDegrees", t)

	tests = []basicTest{{90, 1}, {270, -1}}
	testFunc(CscDegrees, tests, "CscDegrees", t)

	tests = []basicTest{{45, 1}, {-45, -1}}
	testFunc(CotDegrees, tests, "CotDegrees", t)

	tests = []basicTest{{1, 90}, {0, 0}}
	testFunc(ArcsinDegrees, tests, "ArcsinDegrees", t)

	tests = []basicTest{{1, 0}, {0, 90}}
	testFunc(ArccosDegrees, tests, "ArccosDegrees", t)

	tests = []basicTest{{1, 45}, {0, 0}}
	testFunc(ArctanDegrees, tests, "ArctanDegrees", t)

	tests = []basicTest{{1, 0}, {-1, 180}}
	testFunc(ArcsecDegrees, tests, "ArcsecDegrees", t)

	tests = []basicTest{{1, 90}, {-1, -90}}
	testFunc(ArccscDegrees, tests, "ArccscDegrees", t)

	tests = []basicTest{{-1, 135}, {1, 45}}
	testFunc(ArccotDegrees, tests, "ArccotDegrees", t)
}

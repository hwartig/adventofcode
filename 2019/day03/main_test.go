package main

import (
	"reflect"
	"testing"
)

func TestParseWire(t *testing.T) {
	testCases := map[string]Wire{
		"R8":          Wire{[]Step{{"R", 8}}},
		"R8,U5,L5,D3": Wire{[]Step{{"R", 8}, {"U", 5}, {"L", 5}, {"D", 3}}},
	}

	for input, expected := range testCases {
		actual := parseWire(input)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v: Expected: %v but got: %v", input, expected, actual)
		}
	}
}

func TestWirePointCloud(t *testing.T) {
	testCases := map[string]Points{
		"R3":          Points{{1, 0, 1}, {2, 0, 2}, {3, 0, 3}},
		"U3":          Points{{0, 1, 1}, {0, 2, 2}, {0, 3, 3}},
		"L3":          Points{{-1, 0, 1}, {-2, 0, 2}, {-3, 0, 3}},
		"D3":          Points{{0, -1, 1}, {0, -2, 2}, {0, -3, 3}},
		"R1,U1,L2,D2": Points{{-1, 0, 5}, {0, 1, 3}, {1, 0, 1}, {-1, -1, 6}, {-1, 1, 4}, {1, 1, 2}},
	}

	for input, expected := range testCases {
		wire := parseWire(input)
		actual := wire.Points()

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v: Expected: %v but got: %v", input, expected, actual)
		}
	}
}

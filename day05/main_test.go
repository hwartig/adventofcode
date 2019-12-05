package main

import (
	"reflect"
	"testing"
)

func TestSplitToNumbers(t *testing.T) {
	testCases := map[string][]int{
		"1":                   []int{1},
		"1,0,0,0,99":          []int{1, 0, 0, 0, 99},
		"2,3,0,3,99":          []int{2, 3, 0, 3, 99},
		"2,4,4,5,99,0":        []int{2, 4, 4, 5, 99, 0},
		"1,1,1,4,99,5,6,0,99": []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
	}

	for input, expected := range testCases {
		actual := splitToNumbers(input)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v: Expected: %v but got: %v", input, expected, actual)
		}
	}
}

func TestRun(t *testing.T) {
	testCases := map[string][]int{
		// indirect mode
		"1,0,0,0,99":                    []int{2, 0, 0, 0, 99},       // (1 + 1 = 2).
		"2,3,0,3,99":                    []int{2, 3, 0, 6, 99},       // (3 * 2 = 6).
		"2,4,4,5,99,0":                  []int{2, 4, 4, 5, 99, 9801}, // (99 * 99 = 9801).
		"1,1,1,4,99,5,6,0,99":           []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		"1,9,10,3,2,3,11,0,99,30,40,50": []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		"3,0,4,0,99":                    []int{3, 0, 4, 0, 99},
		// direct mode
		"1101,2,3,3,99": []int{1101, 2, 3, 5, 99}, // (2 + 3 = 5).
		"1102,2,3,3,99": []int{1102, 2, 3, 6, 99}, // (2 * 3 = 6).

	}

	for input, expected := range testCases {
		actual := run(input, 0, 0)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("%v: Expected: %v but got: %v", input, expected, actual)
		}
	}
}

func TestParseInstruction(t *testing.T) {
	testCases := map[[4]int]instruction{
		[4]int{99, 0, 0, 0}:    instruction{99, 0, 0, 0, 0, 0, 0},
		[4]int{1, 2, 3, 4}:     instruction{1, 2, 3, 4, 0, 0, 0},
		[4]int{2, 3, 4, 5}:     instruction{2, 3, 4, 5, 0, 0, 0},
		[4]int{12345, 6, 7, 8}: instruction{45, 6, 7, 8, 3, 2, 1},
	}

	for input, expected := range testCases {
		actual := parseInstruction(input[:])

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("%v: Expected: %v but got: %v", input, expected, actual)
		}
	}

	// test that shorter instructions don't throw an error
	actual := parseInstruction([]int{99})
	expected := instruction{99, 0, 0, 0, 0, 0, 0}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("99: Expected: %v but got: %v", expected, actual)
	}
}

package main

import (
	"testing"
)

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

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

		if !sliceEqual(actual, expected) {
			t.Errorf("%v: Expected: %v but got: %v", input, expected, actual)
		}
	}
}

func TestRun(t *testing.T) {
	testCases := map[string][]int{
		"1,0,0,0,99":                    []int{2, 0, 0, 0, 99},       // (1 + 1 = 2).
		"2,3,0,3,99":                    []int{2, 3, 0, 6, 99},       // (3 * 2 = 6).
		"2,4,4,5,99,0":                  []int{2, 4, 4, 5, 99, 9801}, // (99 * 99 = 9801).
		"1,1,1,4,99,5,6,0,99":           []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		"1,9,10,3,2,3,11,0,99,30,40,50": []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
	}

	for input, expectedOutput := range testCases {
		actualOutput := run(input)

		if !sliceEqual(expectedOutput, actualOutput) {
			t.Errorf("%v: Expected: %v but got: %v", input, expectedOutput, actualOutput)
		}
	}
}

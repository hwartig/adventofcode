package main

import (
	"bytes"
	"io"
	"reflect"
	"strings"
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

type RunTestCase struct {
	StartMem       string
	ExpectedEndMem []int
	Input          string
	ExpectedOutput string
}

func (tc RunTestCase) Reader() io.Reader {
	return strings.NewReader(tc.Input)
}

func TestRunInputOutput(t *testing.T) {

	testCases := []RunTestCase{
		// positional mode
		{"1,0,0,0,99", []int{2, 0, 0, 0, 99}, "", ""},         // (1 + 1 = 2).
		{"2,3,0,3,99", []int{2, 3, 0, 6, 99}, "", ""},         // (3 * 2 = 6).
		{"2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}, "", ""}, // (99 * 99 = 9801).
		{"1,1,1,4,99,5,6,0,99", []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, "", ""},
		{"1,9,10,3,2,3,11,0,99,30,40,50", []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, "", ""},

		{"3,5,4,5,99,0", []int{3, 5, 4, 5, 99, 1}, "1", "1"}, // input => output

		{"3,9,8,9,10,9,4,9,99,-1,8", []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8}, "1", "0"},  // equal
		{"3,9,8,9,10,9,4,9,99,-1,8", []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8}, "8", "1"},  // equal
		{"3,9,7,9,10,9,4,9,99,-1,8", []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8}, "1", "1"},  // less than
		{"3,9,7,9,10,9,4,9,99,-1,8", []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8}, "10", "0"}, // less than

		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 1, 1, 9}, "-1", "1"}, // jumps
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9}, "0", "0"},   // jumps

		// immediate mode
		{"1101,2,3,3,99", []int{1101, 2, 3, 5, 99}, "", ""}, // (2 + 3 = 5).
		{"1102,2,3,3,99", []int{1102, 2, 3, 6, 99}, "", ""}, // (2 * 3 = 6).

		{"3,3,1108,-1,8,3,4,3,99", []int{3, 3, 1108, 0, 8, 3, 4, 3, 99}, "1", "0"},  // equal
		{"3,3,1108,-1,8,3,4,3,99", []int{3, 3, 1108, 1, 8, 3, 4, 3, 99}, "8", "1"},  // equal
		{"3,3,1107,-1,8,3,4,3,99", []int{3, 3, 1107, 1, 8, 3, 4, 3, 99}, "1", "1"},  // less than
		{"3,3,1107,-1,8,3,4,3,99", []int{3, 3, 1107, 0, 8, 3, 4, 3, 99}, "10", "0"}, // less than

		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", []int{3, 3, 1105, -7, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, "-7", "1"}, // jumps
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", []int{3, 3, 1105, 0, 9, 1101, 0, 0, 12, 4, 12, 99, 0}, "0", "0"},   // jumps
	}

	for _, tc := range testCases {
		var buf bytes.Buffer
		actual := run(tc.StartMem, tc.Reader(), &buf)

		if !reflect.DeepEqual(tc.ExpectedEndMem, actual) {
			t.Errorf("%v: Expected End Mem differs: %v but got: %v", tc.StartMem, tc.ExpectedEndMem, actual)
		}
		if buf.String() != tc.ExpectedOutput {
			t.Errorf("%v: Expected Output differs: %v but got: %v", tc.StartMem, tc.ExpectedOutput, buf.String())
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

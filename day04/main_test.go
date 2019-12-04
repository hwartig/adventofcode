package main

import "testing"

func TestMeetsCriteria(t *testing.T) {
	testCases := map[int]bool{
		111111: true,
		223450: false,
		123789: false,
		123455: true,
	}

	for input, expected := range testCases {
		actual := meetsCriteria(input)
		if actual != expected {
			t.Errorf("%v: Expected: %v, got: %v", input, expected, actual)
		}
	}
}

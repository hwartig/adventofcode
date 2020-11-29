package main

import "testing"

func TestMeetsCriteriaPart1(t *testing.T) {
	testCases := map[int]bool{
		111111: true,
		223450: false,
		123789: false,
		123455: true,
	}

	for input, expected := range testCases {
		actual := meetsCriteriaPart1(input)
		if actual != expected {
			t.Errorf("%v: Expected: %v, got: %v", input, expected, actual)
		}
	}
}

func TestMeetsCriteriaPart2(t *testing.T) {
	testCases := map[int]bool{
		112233: true,
		123444: false,
		111234: false,
		111122: true,
	}

	for input, expected := range testCases {
		actual := meetsCriteriaPart2(input)
		if actual != expected {
			t.Errorf("%v: Expected: %v, got: %v", input, expected, actual)
		}
	}
}

package main

import "testing"

func TestCalcFuel(t *testing.T) {
	testCases := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}

	for mass, expectedFuel := range testCases {
		actualFuel := CalcFuel(mass)
		if actualFuel != expectedFuel {
			t.Errorf("Fuel requirement for mass %d should be %d but CalcFuel returned %d", mass, expectedFuel, actualFuel)
		}
	}
}

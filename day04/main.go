package main

import (
	"fmt"
	"strconv"
)

func meetsCriteria(i int) bool {
	digits := strconv.Itoa(i)
	hasDoubleDigits := false

	for i := 1; i < len(digits); i++ {
		//Two adjacent digits are the same (like 22 in 122345).
		if digits[i-1] == digits[i] {
			hasDoubleDigits = true
		}

		//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
		if digits[i-1] > digits[i] {
			return false
		}
	}

	return hasDoubleDigits
}

func countValidPasswords(from, to int) int {
	count := 0

	for i := from; i <= to; i++ {
		if meetsCriteria(i) {
			count += 1
		}
	}
	return count
}

func main() {
	fmt.Println("part1: ", countValidPasswords(246540, 787419))
}

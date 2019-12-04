package main

import (
	"fmt"
	"strconv"
)

type Criteria = func(a int) bool

func meetsCriteriaPart2(input int) bool {
	digits := strconv.Itoa(input)
	hasDoubleDigits := false

	for i := 1; i < len(digits); i++ {
		//Two adjacent digits are the same (like 22 in 122345).
		if digits[i-1] == digits[i] {
			// only count group if its not part of a tripple
			if i == 1 && digits[i] != digits[i+1] {
				hasDoubleDigits = true
			}
			if i > 1 && i < len(digits)-1 && digits[i-1] != digits[i-2] && digits[i] != digits[i+1] {
				hasDoubleDigits = true
			}
			if i == len(digits)-1 && digits[i-1] != digits[i-2] {
				hasDoubleDigits = true
			}
		}

		//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
		if digits[i-1] > digits[i] {
			return false
		}
	}

	return hasDoubleDigits
}

func meetsCriteriaPart1(i int) bool {
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

func countValidPasswords(from, to int, c Criteria) int {
	count := 0

	for i := from; i <= to; i++ {
		if c(i) {
			count += 1
		}
	}
	return count
}

func main() {
	//fmt.Println("part1: ", countValidPasswords(246540, 787419, meetsCriteriaPart1))
	fmt.Println("part2: ", countValidPasswords(246540, 787419, meetsCriteriaPart2))
}

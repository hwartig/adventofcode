package main

import (
	"fmt"
	"strconv"

	"../../aoc"
)

func Part1(input string) string {
	result := 0
	for _, c := range input {
		if c == '(' {
			result += 1
		} else if c == ')' {
			result -= 1
		} else {
			//	log.Fatal("found unknown rune", c)
		}
	}
	return strconv.Itoa(result)
}

func Part2(input string) string {
	result := 0
	for i, c := range input {
		if c == '(' {
			result += 1
		} else if c == ')' {
			result -= 1
		}

		if result == -1 {
			return strconv.Itoa(i + 1)
		}
	}
	return strconv.Itoa(-1)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 138
	fmt.Println(Part2(aoc.ReadInput())) // 1771
}

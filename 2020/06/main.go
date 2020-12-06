package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func Tally(input string) map[rune]int {
	tally := make(map[rune]int)
	for _, c := range input {
		if c == '\n' {
			continue
		}
		tally[c] += 1
	}
	return tally
}

func Part1(input string) string {
	result := 0
	groups := strings.Split(input, "\n\n")

	for _, g := range groups {
		result += len(Tally(g))
	}

	return strconv.Itoa(result)
}

func Part2(input string) string {
	result := 0

	groups := strings.Split(input, "\n\n")

	for _, g := range groups {
		tallies := Tally(g)

		persons := strings.Count(g, "\n") + 1

		for _, v := range tallies {
			if v == persons {
				result += 1
			}
		}
	}

	return strconv.Itoa(result)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 7027
	fmt.Println(Part2(aoc.ReadInput())) // 3579
}

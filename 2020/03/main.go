package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func CountTrees(lines []string, right, down int) int {
	result := 0
	width := len(lines[0])

	for i, j := 0, 0; i < len(lines); i, j = i+down, (j+right)%width {
		if string(lines[i][j]) == "#" {
			result += 1
		}
	}

	return result
}

func Part1(input string) string {
	lines := strings.Split(input, "\n")

	return strconv.Itoa(CountTrees(lines, 3, 1))
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")

	result := CountTrees(lines, 1, 1) *
		CountTrees(lines, 3, 1) *
		CountTrees(lines, 5, 1) *
		CountTrees(lines, 7, 1) *
		CountTrees(lines, 1, 2)

	return strconv.Itoa(result)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 230
	fmt.Println(Part2(aoc.ReadInput())) // 9533698720
}

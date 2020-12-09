package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"../../aoc"
)

var preambleLength = 25

func ValidPart1(pre []int, num int) bool {
	sorted := append([]int(nil), pre...)

	sort.Ints(sorted)

	if num < sorted[0]+sorted[1] || num > sorted[preambleLength-2]+sorted[preambleLength-1] {
		return false
	}

	for i := 0; i < len(sorted); i++ {
		for j := 0; j < len(sorted); j++ {
			if sorted[i]+sorted[j] == num {
				return true
			}
		}
	}

	return false
}

func Parse(input string) (nums []int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		nums = append(nums, aoc.Atoi(line))
	}

	return nums
}

func Part1(input string) string {
	nums := Parse(input)

	for i := 0; i < len(nums)-preambleLength; i++ {
		if !ValidPart1(nums[i:i+preambleLength], nums[i+preambleLength]) {
			return strconv.Itoa(nums[i+preambleLength])
		}
	}

	return "no invalid number found"
}

func ValidPart2(nums []int, sum int) bool {
	return false
}

func Part2(input string) string {
	goal := aoc.Atoi(Part1(input)) // 1639024365

	nums := Parse(input)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if aoc.Sum(nums[i:j+1]) == goal {
				min, max := aoc.MinMax(nums[i : j+1])

				return strconv.Itoa(min + max)
			}
		}
	}

	return "didn't find a solution for Part 2"
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 1639024365
	fmt.Println(Part2(aoc.ReadInput())) // 219202240
}

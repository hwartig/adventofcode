package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func Solve(input string, iterations int) int {
	numMem1 := make(map[int]int) // most recently spoken
	numMem2 := make(map[int]int) // spoken before that

	nums := strings.Split(input, ",")
	lastNumSpoken := 0

	// this only works as I know that there are no duplicates in input
	for i, n := range nums {
		lastNumSpoken = aoc.Atoi(n)
		numMem1[lastNumSpoken] = i + 1
	}

	for i := len(nums) + 1; i <= iterations; i++ {

		if k, ok := numMem2[lastNumSpoken]; ok {
			lastNumSpoken = numMem1[lastNumSpoken] - k
		} else {
			lastNumSpoken = 0
		}

		if _, ok := numMem1[lastNumSpoken]; ok {
			numMem2[lastNumSpoken] = numMem1[lastNumSpoken]
		}

		numMem1[lastNumSpoken] = i
	}

	return lastNumSpoken
}

func Part1(input string) string {
	return strconv.Itoa(Solve(input, 2020))
}

func Part2(input string) string {
	return strconv.Itoa(Solve(input, 30000000))
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 376
	fmt.Println(Part2(aoc.ReadInput())) // 323780
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func Solve(input string, iterations int) int {
	//fmt.Printf("solving %s for %d iterations\n", input, iterations)
	numMem1 := make(map[int]int) // most recently spoken
	numMem2 := make(map[int]int) // spoken before that

	nums := strings.Split(input, ",")
	lastNumSpoken := 0

	// starting nums
	for i, n := range nums {
		lastNumSpoken = aoc.Atoi(n)
		numMem1[lastNumSpoken] = i + 1
	}

	//fmt.Println(numMem1, numMem2)

	// knowing that there no duplicates in input
	for i := len(nums) + 1; i <= iterations; i++ {
		//fmt.Print("i= ", i, ": last: ", lastNumSpoken, numMem1, numMem2)

		if k, ok := numMem2[lastNumSpoken]; ok {
			lastNumSpoken = numMem1[lastNumSpoken] - k
		} else {
			lastNumSpoken = 0
		}

		if _, ok := numMem1[lastNumSpoken]; ok {
			numMem2[lastNumSpoken] = numMem1[lastNumSpoken]
		}

		numMem1[lastNumSpoken] = i

		//fmt.Println(" ==> last:", lastNumSpoken, numMem1, numMem2)
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
	aoc.AssertEq(Solve("0,3,6", 4), 0)
	aoc.AssertEq(Solve("0,3,6", 5), 3)
	aoc.AssertEq(Solve("0,3,6", 6), 3)
	aoc.AssertEq(Solve("0,3,6", 7), 1)
	aoc.AssertEq(Solve("0,3,6", 8), 0)
	aoc.AssertEq(Solve("0,3,6", 9), 4)
	aoc.AssertEq(Solve("0,3,6", 10), 0)

	aoc.AssertEq(Solve("1,3,2", 2020), 1)
	aoc.AssertEq(Solve("2,1,3", 2020), 10)
	aoc.AssertEq(Solve("1,2,3", 2020), 27)
	aoc.AssertEq(Solve("2,3,1", 2020), 78)
	aoc.AssertEq(Solve("3,2,1", 2020), 438)
	aoc.AssertEq(Solve("3,1,2", 2020), 1836)

	fmt.Println(Part1(aoc.ReadInput())) // 376
	fmt.Println(Part2(aoc.ReadInput())) // 323780
}

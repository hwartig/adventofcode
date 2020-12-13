package main

import (
	"fmt"
	"log"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"../../aoc"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")

	earliest := aoc.Atoi(lines[0])

	var pairs []aoc.Pair

	for _, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			continue
		}

		busNum := aoc.Atoi(bus)

		if !big.NewInt(int64(busNum)).ProbablyPrime(0) {
			log.Fatalf("%v is not a prime", bus)
		}

		for departsAt := earliest; departsAt < earliest+busNum; departsAt++ {
			if departsAt%busNum == 0 {
				pairs = append(pairs, aoc.Pair{busNum, departsAt - earliest})
			}
		}
	}

	sort.Sort(aoc.ByValue(pairs))

	return strconv.Itoa(pairs[0].Key * (pairs[0].Val))
}

func MinimalTimestamp(input string) (pos int) {
	var pairs []aoc.Pair

	for i, bus := range strings.Split(input, ",") {
		if bus == "x" {
			continue
		}
		busNum := aoc.Atoi(bus)

		pairs = append(pairs, aoc.Pair{busNum, i})
	}

	inc := pairs[0].Key // we can never be earlier than that

	for _, bus := range pairs[1:len(pairs)] {
		for {
			pos += inc
			// add increment until we find a valid timestamp
			if (pos+bus.Val)%bus.Key == 0 {
				break
			}
		}
		inc *= bus.Key
	}

	return pos
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")

	return strconv.Itoa(MinimalTimestamp(lines[1]))
}

func main() {
	aoc.AssertEq(Part1(`939
7,13,x,x,59,x,31,19`), "295")

	aoc.AssertEq(MinimalTimestamp("2,3"), 2)
	aoc.AssertEq(MinimalTimestamp("2,x,3"), 4)
	aoc.AssertEq(MinimalTimestamp("7,13,x,x,59,x,31,19"), 1068781)
	aoc.AssertEq(MinimalTimestamp("17,x,13,19"), 3417)
	aoc.AssertEq(MinimalTimestamp("67,7,59,61"), 754018)
	aoc.AssertEq(MinimalTimestamp("67,x,7,59,61"), 779210)
	aoc.AssertEq(MinimalTimestamp("67,7,x,59,61"), 1261476)
	aoc.AssertEq(MinimalTimestamp("1789,37,47,1889"), 1202161486)

	fmt.Println(Part1(aoc.ReadInput())) // 5946
	fmt.Println(Part2(aoc.ReadInput())) // 645338524823718
}

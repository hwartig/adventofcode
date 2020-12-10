package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"../../aoc"
)

func Parse(input string) (nums []int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		nums = append(nums, aoc.Atoi(line))
	}

	sort.Ints(nums)
	return nums
}

func Part1(input string) string {
	nums := Parse(input)

	prev, one, three := 0, 0, 1

	for _, n := range nums {
		diff := n - prev
		if diff == 1 {
			one += 1
		} else if diff == 3 {
			three += 1
		}
		prev = n
	}

	return strconv.Itoa(one * three)
}

type key struct {
	Start, Len int
}

var seen = make(map[key]int)

func Count(adapters []int, start, end int) int {
	k := key{start, len(adapters)}
	if val, ok := seen[k]; ok {
		return val
	}

	result := 0

	if end-start <= 3 {
		result += 1
	}

	if len(adapters) < 1 {
		return result
	}

	if adapters[0]-start <= 3 {
		result += Count(adapters[1:len(adapters)], adapters[0], end)
	}

	result += Count(adapters[1:len(adapters)], start, end)

	seen[k] = result

	return result
}

func Part2(input string) string {
	nums := Parse(input)

	return strconv.Itoa(Count(nums, 0, aoc.Max(nums)+3))
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 1998
	fmt.Println(Part2(aoc.ReadInput())) // 347250213298688
}

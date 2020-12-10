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

type IntSlice []int

var seen = make(map[int]int)

func (nums IntSlice) Count(i int) int {
	result := 0

	if i == len(nums)-1 {
		return 1
	}

	if count, ok := seen[i]; ok {
		return count
	}

	for j := i + 1; j < i+4 && j < len(nums); j++ {
		if nums[j]-nums[i] <= 3 {
			result += nums.Count(j)
		}
	}
	seen[i] = result

	return result
}

func Part2(input string) string {
	nums := append([]int{0}, Parse(input)...)
	nums = append(nums, aoc.Max(nums)+3)

	return strconv.Itoa(IntSlice(nums).Count(0))
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 1998
	fmt.Println(Part2(aoc.ReadInput())) // 347250213298688
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func Parse(input string) (lower, upper int, letter, pwd string) {
	parts := strings.Split(input, " ")
	bounds := strings.Split(parts[0], "-")
	lower, upper = aoc.Atoi(bounds[0]), aoc.Atoi(bounds[1])
	letter = string(parts[1][0])
	pwd = parts[2]
	return
}

func ValidPart1(input string) bool {
	lower, upper, letter, pwd := Parse(input)
	count := strings.Count(pwd, letter)

	return count >= lower && count <= upper
}

func Part1(input string) string {
	result := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if ValidPart1(line) {
			result += 1
		}
	}
	return strconv.Itoa(result)
}

func ValidPart2(input string) bool {
	pos1, pos2, letter, pwd := Parse(input)

	return (pwd[pos1-1] == letter[0]) != (pwd[pos2-1] == letter[0])
}

func Part2(input string) string {
	result := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if ValidPart2(line) {
			result += 1
		}
	}
	return strconv.Itoa(result)
}

func main() {
	//fmt.Println(Part2("1-3 a: abcde"))
	//fmt.Println(Part2("1-3 b: cdefg"))
	//fmt.Println(Part2("2-9 c: ccccccccc"))
	//fmt.Println(Part2("2-9 c: cccccccccc"))
	fmt.Println(Part1(aoc.ReadInput())) // 393
	fmt.Println(Part2(aoc.ReadInput())) // 690
}

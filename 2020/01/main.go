package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func Part1(input string) string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		num1 := aoc.Atoi(line)
		for j := i + 1; j < len(lines); j++ {
			num2 := aoc.Atoi(lines[j])

			if num1+num2 == 2020 {
				return strconv.Itoa(num1 * num2)
			}
		}
	}
	return ""
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		num1 := aoc.Atoi(line)

		for j := i + 1; j < len(lines); j++ {
			num2 := aoc.Atoi(lines[j])

			for k := j; k < len(lines); k++ {
				num3 := aoc.Atoi(lines[k])

				if num1+num2+num3 == 2020 {
					return strconv.Itoa(num1 * num2 * num3)
				}
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 1018944
	fmt.Println(Part2(aoc.ReadInput())) // 8446464
}

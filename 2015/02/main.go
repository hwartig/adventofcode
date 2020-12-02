package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"../../aoc"
)

type Line struct {
	Length int
	Width  int
	Height int
}

func (l Line) Paper() int {
	a := []int{2 * l.Length * l.Width, 2 * l.Width * l.Height, 2 * l.Height * l.Length}
	return aoc.Sum(a) + aoc.Min(a)/2
}

func (l Line) Ribbon() int {
	dims := []int{l.Length, l.Width, l.Height}
	sort.Ints(dims)
	return 2*dims[0] + 2*dims[1] + l.Length*l.Width*l.Height
}

func ParseLine(input string) Line {
	d := strings.Split(input, "x")

	return Line{
		aoc.Atoi(d[0]),
		aoc.Atoi(d[1]),
		aoc.Atoi(d[2]),
	}
}

func Part1(input string) string {
	lines := strings.Split(input, "\n")
	total := 0

	for _, lineString := range lines {
		if len(lineString) > 0 {
			total += ParseLine(lineString).Paper()
		}
	}

	return strconv.Itoa(total)
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")
	total := 0

	for _, lineString := range lines {
		if len(lineString) > 0 {
			total += ParseLine(lineString).Ribbon()
		}
	}

	return strconv.Itoa(total)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 1588178
	fmt.Println(Part2(aoc.ReadInput())) // 3783758
}

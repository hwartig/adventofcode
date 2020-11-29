package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Min(a []int) int {
	min, _ := MinMax(a)
	return min
}

func Max(a []int) int {
	_, max := MinMax(a)
	return max
}

func Sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

type Line struct {
	Length int
	Width  int
	Height int
}

func (l Line) Paper() int {
	a := []int{2 * l.Length * l.Width, 2 * l.Width * l.Height, 2 * l.Height * l.Length}
	return Sum(a) + Min(a)/2
}

func (l Line) Ribbon() int {
	dims := []int{l.Length, l.Width, l.Height}
	sort.Ints(dims)
	return 2*dims[0] + 2*dims[1] + l.Length*l.Width*l.Height
}

func ParseLine(input string) Line {
	d := strings.Split(input, "x")

	return Line{
		Atoi(d[0]),
		Atoi(d[1]),
		Atoi(d[2]),
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

func ReadInput() string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func main() {
	//fmt.Println(Part1(ReadInput()))
	fmt.Println(Part2(ReadInput()))
}

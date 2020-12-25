package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

const remainder = 20201227

func TransformSubjectNumber(num, loopSize int) int {
	result := 1

	for i := 0; i < loopSize; i++ {
		result = (result * num) % remainder
	}

	return result
}

func Parse(input string) (int, int) {
	lines := strings.Split(input, "\n")
	return aoc.Atoi(lines[0]), aoc.Atoi(lines[1])
}

func GuessLoopSize(pubKey int) int {
	result := 1

	for i := 0; i < 100000000; i++ {
		result = (result * 7) % remainder
		if pubKey == result {
			return i + 1
		}
	}

	return 0
}

func Part1(input string) string {
	pubCard, pubDoor := Parse(input)

	loopCard := GuessLoopSize(pubCard)

	encDoor := TransformSubjectNumber(pubDoor, loopCard)

	return strconv.Itoa(encDoor)
}

func main() {
	sample := `17807724
5764801`
	fmt.Println(Part1(sample))
	fmt.Println(Part1(aoc.ReadInput()))
}

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../aoc"
)

func ApplyMask(value int, mask string) int {
	maskInt, err := strconv.ParseInt(strings.ReplaceAll(mask, "X", "0"), 2, 0)
	if err != nil {
		log.Panic(err)
	}
	value |= int(maskInt)

	maskInt, err = strconv.ParseInt(strings.ReplaceAll(mask, "X", "1"), 2, 0)
	if err != nil {
		log.Panic(err)
	}

	value &= int(maskInt)
	return value
}

func SumMemValues(mem map[int]int) (result int) {
	for _, val := range mem {
		result += val
	}

	return result
}

func Part1(input string) string {
	mem := make(map[int]int)
	mask := ""
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		op := parts[0]
		val := parts[1]
		if op == "mask" {
			mask = val
		} else {
			pos := aoc.Atoi(op[4 : len(op)-1])
			mem[pos] = ApplyMask(aoc.Atoi(val), mask)
		}
	}

	return strconv.Itoa(SumMemValues(mem))
}

func GenAllPos(pos int, mask string) (result []int) {
	if mask == "" {
		result = []int{0}
	} else {

		for _, m := range GenAllPos(pos/2, mask[0:len(mask)-1]) {
			bit := mask[len(mask)-1]

			if bit == '0' {
				result = append(result, 2*m+pos%2) // unchanged
			} else if bit == '1' {
				result = append(result, 2*m+1) // set to 1
			} else if bit == 'X' {
				result = append(result, 2*m+0)
				result = append(result, 2*m+1)
			}
		}
	}
	return result
}

func Part2(input string) string {
	mem := make(map[int]int)
	mask := ""
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		op := parts[0]
		val := parts[1]
		if op == "mask" {
			mask = val
		} else {
			pos := aoc.Atoi(op[4 : len(op)-1])
			valInt := aoc.Atoi(val)

			for _, p := range GenAllPos(pos, mask) {
				mem[p] = valInt
			}
		}
	}
	return strconv.Itoa(SumMemValues(mem))
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 11179633149677
	fmt.Println(Part2(aoc.ReadInput())) // 4822600194774
}

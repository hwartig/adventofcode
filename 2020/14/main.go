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

type Memory map[int]int

func (m Memory) SumMemValues() (result int) {
	for _, val := range m {
		result += val
	}

	return result
}

func (m Memory) Run(input string, apply func(int, int, string)) {
	mask := ""
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		op := parts[0]
		val := parts[1]
		if op == "mask" {
			mask = val
		} else {
			apply(aoc.Atoi(op[4:len(op)-1]), aoc.Atoi(val), mask)
		}
	}
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
			} else {
				log.Panicf("don't know what to do with bit %s", bit)
			}
		}
	}
	return result
}

func Part1(input string) string {
	mem := make(Memory)

	mem.Run(input, func(pos, val int, mask string) {
		mem[pos] = ApplyMask(val, mask)
	})

	return strconv.Itoa(mem.SumMemValues())
}

func Part2(input string) string {
	mem := make(Memory)

	mem.Run(input, func(pos, val int, mask string) {
		for _, p := range GenAllPos(pos, mask) {
			mem[p] = val
		}
	})

	return strconv.Itoa(mem.SumMemValues())
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 11179633149677
	fmt.Println(Part2(aoc.ReadInput())) // 4822600194774
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

type Instruction struct {
	Op  string
	Val int
}

type Program []Instruction

func Parse(input string) (insts Program) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		insts = append(insts, ParseLine(line))
	}

	return insts
}

func ParseLine(line string) Instruction {
	parts := strings.Split(line, " ")

	return Instruction{Op: parts[0], Val: aoc.Atoi(parts[1])}
}

func Part1(input string) string {
	insts := Parse(input)

	acc, _ := insts.Run()

	return strconv.Itoa(acc)
}

func (insts Program) Run() (int, bool) {
	seen := make(map[int]bool)
	acc := 0

	for i := 0; i < len(insts); i++ {
		if seen[i] {
			return acc, false
		}

		seen[i] = true

		switch insts[i].Op {
		case "acc":
			acc += insts[i].Val
		case "jmp":
			i += insts[i].Val - 1
			continue
		case "nop":
		}
	}

	return acc, true
}

func Part2(input string) string {
	insts := Parse(input)

	for i, inst := range insts {
		switch inst.Op {
		case "nop":
			insts[i] = Instruction{"jmp", inst.Val}
		case "jmp":
			insts[i] = Instruction{"nop", inst.Val}
		}

		if acc, ok := insts.Run(); ok {
			return strconv.Itoa(acc)
		}
		insts[i] = inst

	}

	return "Program didn't terminate"
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 1753
	fmt.Println(Part2(aoc.ReadInput())) // 733
}

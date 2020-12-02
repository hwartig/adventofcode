package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../aoc"
)

type Action int

const (
	TurnOn Action = iota
	TurnOff
	Toggle
)

type Instruction struct {
	Action                 Action
	FromX, FromY, ToX, ToY int
}

func ParseLine(line string) Instruction {
	var action Action
	var rest string

	if strings.HasPrefix(line, "toggle ") {
		action = Toggle
		rest = line[len("toggle "):]
	} else if strings.HasPrefix(line, "turn off ") {
		action = TurnOff
		rest = line[len("turn off "):]
	} else if strings.HasPrefix(line, "turn on ") {
		action = TurnOn
		rest = line[len("turn on "):]
	} else {
		log.Fatalf("Unkown action in '%s'\n", line)
	}

	coords := strings.Split(rest, " through ")
	from, to := strings.Split(coords[0], ","), strings.Split(coords[1], ",")

	return Instruction{
		action,
		aoc.Atoi(from[0]),
		aoc.Atoi(from[1]),
		aoc.Atoi(to[0]),
		aoc.Atoi(to[1]),
	}
}

func Parse(input string) []Instruction {
	var instructions []Instruction

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		instructions = append(instructions, ParseLine(line))
	}

	return instructions
}

type BoolGrid [1000][1000]bool

func (g *BoolGrid) Apply(i Instruction) {
	for x := i.FromX; x <= i.ToX; x++ {
		for y := i.FromY; y <= i.ToY; y++ {
			switch i.Action {
			case TurnOn:
				g[x][y] = true
			case TurnOff:
				g[x][y] = false
			case Toggle:
				g[x][y] = !g[x][y]
			}
		}
	}
}

func (g *BoolGrid) Count() int {
	count := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] {
				count += 1
			}
		}
	}

	return count
}

func Part1(input string) string {
	var grid BoolGrid
	instructions := Parse(input)
	for _, inst := range instructions {
		grid.Apply(inst)
	}

	return strconv.Itoa(grid.Count())
}

type IntGrid [1000][1000]int

func (g *IntGrid) Apply(i Instruction) {
	for x := i.FromX; x <= i.ToX; x++ {
		for y := i.FromY; y <= i.ToY; y++ {
			switch i.Action {
			case TurnOn:
				g[x][y] += 1
			case TurnOff:
				if g[x][y] > 0 {
					g[x][y] -= 1
				}
			case Toggle:
				g[x][y] += 2
			}
		}
	}
}

func (g *IntGrid) Count() int {
	count := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			count += g[i][j]
		}
	}

	return count
}

func Part2(input string) string {
	var grid IntGrid
	instructions := Parse(input)
	for _, inst := range instructions {
		grid.Apply(inst)
	}

	return strconv.Itoa(grid.Count())
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 377891
	fmt.Println(Part2(aoc.ReadInput())) // 14110788
}

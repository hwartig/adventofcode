package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

type Instruction struct {
	Dir  rune // direction
	Dist int  // distance
}

func Parse(input string) (res []Instruction) {
	for _, l := range strings.Split(input, "\n") {
		res = append(res, Instruction{rune(l[0]), aoc.Atoi(l[1:len(l)])})
	}
	return
}

func Part1(input string) string {
	insts := Parse(input)
	shipX := 0
	shipY := 0
	dir := 90 // north is up at 0, increase clockwise, so our start east is 90
	dirs := []rune("NESW")

	for _, inst := range insts {
		action := inst.Dir

		if inst.Dir == 'F' {
			action = dirs[dir/90]
		}

		if action == 'N' {
			shipY += inst.Dist
		}
		if action == 'S' {
			shipY -= inst.Dist
		}
		if action == 'W' {
			shipX -= inst.Dist
		}
		if action == 'E' {
			shipX += inst.Dist
		}

		if inst.Dir == 'L' {
			dir = (360 + dir - inst.Dist) % 360
		}
		if inst.Dir == 'R' {
			dir = (dir + inst.Dist) % 360
		}
	}

	return strconv.Itoa(aoc.Abs(shipY) + aoc.Abs(shipX))
}

func Part2(input string) string {
	insts := Parse(input)
	wpX := 10
	wpY := 1
	shipX := 0
	shipY := 0

	for _, inst := range insts {
		// move waypoint
		if inst.Dir == 'N' {
			wpY += inst.Dist
		}
		if inst.Dir == 'S' {
			wpY -= inst.Dist
		}
		if inst.Dir == 'W' {
			wpX -= inst.Dist
		}
		if inst.Dir == 'E' {
			wpX += inst.Dist
		}

		// move ship
		if inst.Dir == 'F' {
			shipX += wpX * inst.Dist
			shipY += wpY * inst.Dist
		}

		if inst.Dir == 'L' {
			for i := inst.Dist; i > 0; i -= 90 {
				wpX, wpY = -wpY, wpX
			}
		}
		if inst.Dir == 'R' {
			for i := inst.Dist; i > 0; i -= 90 {
				wpX, wpY = wpY, -wpX
			}
		}
	}

	return strconv.Itoa(aoc.Abs(shipY) + aoc.Abs(shipX))
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 759
	fmt.Println(Part2(aoc.ReadInput())) // 45763
}

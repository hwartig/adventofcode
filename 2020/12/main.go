package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

type Inst struct {
	Dir  rune
	Dist int
}

func Parse(input string) (res []Inst) {
	for _, l := range strings.Split(input, "\n") {
		res = append(res, Inst{rune(l[0]), aoc.Atoi(l[1:len(l)])})
	}
	return
}

func Part1(input string) string {
	insts := Parse(input)
	eastWest := 0
	northSouth := 0
	dir := 'E'
	dirs := "NWSE"

	for _, inst := range insts {
		if inst.Dir == 'N' {
			northSouth += inst.Dist
		}
		if inst.Dir == 'S' {
			northSouth -= inst.Dist
		}
		if inst.Dir == 'W' {
			eastWest -= inst.Dist
		}
		if inst.Dir == 'E' {
			eastWest += inst.Dist
		}

		if inst.Dir == 'F' {
			if dir == 'N' {
				northSouth += inst.Dist
			}
			if dir == 'S' {
				northSouth -= inst.Dist
			}
			if dir == 'W' {
				eastWest -= inst.Dist
			}
			if dir == 'E' {
				eastWest += inst.Dist
			}
		}

		if inst.Dir == 'L' {
			di := inst.Dist / 90
			i := strings.IndexRune(dirs, dir)
			dir = rune(dirs[(i+di)%4])
		}
		if inst.Dir == 'R' {
			di := inst.Dist / 90
			i := strings.IndexRune(dirs, dir)
			dir = rune(dirs[(4+i-di)%4])
		}
	}

	return strconv.Itoa(aoc.Abs(northSouth) + aoc.Abs(eastWest))
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
	fmt.Println(Part2(`F10
N3
F7
R90
F11`))
	fmt.Println(Part1(aoc.ReadInput()))
	fmt.Println(Part2(aoc.ReadInput()))
}

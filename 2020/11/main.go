package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

type Grid [][]rune

func Parse(input string) (grid Grid) {
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune("."+line+"."))
	}

	return
}

func CountOccupiedAdjacent(grid Grid, r, c int) int {
	result := 0

	for _, dr := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			if (dr == 0 && dc == 0) || r+dr < 0 || r+dr == len(grid) || c+dc < 0 || c+dc == len(grid[0]) {
				continue
			}

			if grid[r+dr][c+dc] == '#' {
				result += 1
			}
		}
	}

	return result
}

func CountOccupiedVisible(grid Grid, r, c int) int {
	result := 0

	for _, dr := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			if dr == 0 && dc == 0 {
				continue
			}

			for i := 1; 0 <= r+i*dr && r+i*dr < len(grid) && 0 <= c+i*dc && c+i*dc < len(grid[0]); i++ {
				ch := grid[r+i*dr][c+i*dc]
				if ch == '#' {
					result += 1
					break
				} else if ch == 'L' {
					break
				}
			}
		}
	}

	return result
}

func (grid Grid) Copy() (copy Grid) {
	for _, line := range grid {
		copy = append(copy, append([]rune{}, line...))
	}
	return
}

func (grid Grid) Step(max int, count func(Grid, int, int) int) (Grid, int) {
	changes := 0
	copy := grid.Copy()

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			c := grid[i][j]
			if c == 'L' && count(grid, i, j) == 0 {
				copy[i][j] = '#'
				changes += 1
			} else if c == '#' && count(grid, i, j) >= max {
				copy[i][j] = 'L'
				changes += 1
			}
		}
	}

	return copy, changes
}

func (grid Grid) String() (result string) {
	for _, line := range grid {
		result += string(line) + "\n"
	}
	return result
}

func (grid Grid) CountOccupiedSeats() (result int) {
	for _, line := range grid {
		for _, seat := range line {
			if seat == '#' {
				result += 1
			}
		}
	}

	return result
}

func (grid Grid) Run(max int, count func(Grid, int, int) int) string {
	for i := 0; i < 1000; i++ {
		next, changes := grid.Step(max, count)

		//fmt.Println(grid)

		if changes == 0 {
			return strconv.Itoa(grid.CountOccupiedSeats())
		}

		grid = next
	}
	return "no solution found in 1000 steps"
}

func Part1(input string) string {
	grid := Parse(input)

	return grid.Run(4, CountOccupiedAdjacent)
}

func Part2(input string) string {
	grid := Parse(input)

	return grid.Run(5, CountOccupiedVisible)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 2453
	fmt.Println(Part2(aoc.ReadInput())) // 2159
}

package main

import (
	"fmt"
	"strconv"

	"../../aoc"
)

func Part1(input string) string {
	seen := make(map[string]bool)
	seen["0,0"] = true
	y, x := 0, 0

	for _, c := range input {
		if c == '<' {
			x -= 1
		} else if c == '>' {
			x += 1
		} else if c == '^' {
			y += 1
		} else if c == 'v' {
			y -= 1
		}

		seen[fmt.Sprintf("%d,%d", x, y)] = true
	}

	return strconv.Itoa(len(seen))
}

func Part2(input string) string {
	seen := make(map[string]bool)
	seen["0,0"] = true
	sy, sx := 0, 0
	ry, rx := 0, 0

	for i, c := range input {
		if i%2 == 0 {
			if c == '<' {
				sx -= 1
			} else if c == '>' {
				sx += 1
			} else if c == '^' {
				sy += 1
			} else if c == 'v' {
				sy -= 1
			}
			seen[fmt.Sprintf("%d,%d", sx, sy)] = true
		} else {
			if c == '<' {
				rx -= 1
			} else if c == '>' {
				rx += 1
			} else if c == '^' {
				ry += 1
			} else if c == 'v' {
				ry -= 1
			}
		}
		seen[fmt.Sprintf("%d,%d", rx, ry)] = true
	}

	return strconv.Itoa(len(seen))
}

func main() {
	//fmt.Println(Part2("^v"))
	//fmt.Println(Part2("^>v<"))
	//fmt.Println(Part2("^v^v^v^v^v"))
	fmt.Println(Part1(aoc.ReadInput())) // 2572
	fmt.Println(Part2(aoc.ReadInput())) // 2631
}

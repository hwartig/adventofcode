package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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
	//fmt.Println(Part2("^v"))
	//fmt.Println(Part2("^>v<"))
	//fmt.Println(Part2("^v^v^v^v^v"))
}

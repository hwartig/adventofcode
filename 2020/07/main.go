package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"../../aoc"
)

//var instructionRe = regexp.MustCompile("^(\\w+ \\w+) bags contain( (\\d)? ?(\\w+ \\w+) bags?(,|.))+$")
//var instructionRe = regexp.MustCompile("^(\\w+ \\w+) bags contain( \\d? ?\\w+ \\w+ bags?(,|.))+$")

var containedin = make(map[string]map[string]int)
var contains = make(map[string]map[string]int)

func Parse(line string) {
	parts := strings.Split(line, " contain")
	outerColor := parts[0][0 : len(parts[0])-len(" bags")]

	inner := strings.Split(parts[1][0:len(parts[1])-1], ",")

	//fmt.Println(outerColor, ":")
	for _, a := range inner {
		innerParts := strings.Split(a[1:], " ")
		if len(innerParts) < 4 {
			// no other bags case
		} else if len(innerParts) == 4 {
			num := aoc.Atoi(innerParts[0])
			innerColor := innerParts[1] + " " + innerParts[2]

			//fmt.Println("-", outerColor, "->", num, ":", innerColor)

			colors, ok := containedin[innerColor]
			if !ok {
				colors = make(map[string]int)
				containedin[innerColor] = colors
			}
			colors[outerColor] = 1

			colors, ok = contains[outerColor]
			if !ok {
				colors = make(map[string]int)
				contains[outerColor] = colors
			}
			colors[innerColor] += num
		} else {
			log.Fatalf("error couldn not parse '%d'", a)
		}
	}
}

var holdsgold = make(map[string]bool)

func Check(color string) int {
	for k := range containedin[color] {
		holdsgold[k] = true
		Check(k)
	}
	return len(holdsgold)
}

func Part1(input string) string {
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		Parse(l)
	}

	return strconv.Itoa(Check("shiny gold"))
}

func Cost(color string) int {
	result := 0
	for k, v := range contains[color] {
		result += v
		result += v * Cost(k)
	}
	return result
}

func Part2(input string) string {
	return strconv.Itoa(Cost("shiny gold"))
}

func main() {
	fmt.Println(Part1(aoc.ReadInput()))
	fmt.Println(Part2(""))
}

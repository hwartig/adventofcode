package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"../../aoc"
)

type Ops int

var binary = regexp.MustCompile("^(\\w+) (RSHIFT|LSHIFT|AND|OR) (\\w+) -> (\\w{1,2})$")
var not = regexp.MustCompile("^NOT (\\w+) -> (\\w{1,2})$")
var lineIn = regexp.MustCompile("^(\\w+) -> (\\w{1,2})$")

const (
	NOT Ops = iota
	RSHIFT
	OR
	AND
	LSHIFT
)

type Gate struct {
	In1, In2, Op, Out string
}

func (g *Gate) HasAllInputs(state map[string]int) (bool, int, int) {
	ok1, ok2 := true, true

	val1, err := strconv.Atoi(g.In1)
	if err != nil {
		val1, ok1 = state[g.In1]
	}

	val2, err := strconv.Atoi(g.In2)
	if err != nil {
		val2, ok2 = state[g.In2]
	}

	return ok1 && ok2, val1, val2
}

func (g *Gate) Apply(state map[string]int) bool {
	ok, in1, in2 := g.HasAllInputs(state)
	if !ok {
		return false
	}

	switch g.Op {
	case "":
		state[g.Out] = in1
	case "NOT":
		state[g.Out] = ^in1
	case "OR":
		state[g.Out] = in1 | in2
	case "AND":
		state[g.Out] = in1 & in2
	case "LSHIFT":
		state[g.Out] = in1 << in2
	case "RSHIFT":
		state[g.Out] = in1 >> in2
	}

	return true
}

func Parse(line string) Gate {
	matches := binary.FindAllStringSubmatch(line, -1)
	if len(matches) > 0 {
		return Gate{
			In1: matches[0][1],
			Op:  matches[0][2],
			In2: matches[0][3],
			Out: matches[0][4],
		}
	}

	matches = not.FindAllStringSubmatch(line, -1)
	if len(matches) > 0 {
		return Gate{
			In1: matches[0][1],
			Op:  "NOT",
			In2: matches[0][1],
			Out: matches[0][2],
		}
	}

	matches = lineIn.FindAllStringSubmatch(line, -1)
	if len(matches) > 0 {
		return Gate{
			In1: matches[0][1],
			Op:  "",
			In2: matches[0][1],
			Out: matches[0][2],
		}
	}

	log.Fatalf("could not match %s", line)
	return Gate{}
}

func ApplyAll(gates []Gate, state map[string]int) []Gate {
	var result []Gate

	for _, g := range gates {
		if !g.Apply(state) {
			result = append(result, g)
		}
	}

	return result
}

func Part1(input string) string {
	lines := strings.Split(input, "\n")

	var gates []Gate

	for _, line := range lines {
		gates = append(gates, Parse(line))
	}

	state := make(map[string]int)

	for len(gates) > 0 {
		gates = ApplyAll(gates, state)
	}

	return strconv.Itoa(state["a"])
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")

	var gates []Gate

	for _, line := range lines {
		g := Parse(line)

		if g.Out == "b" {
			override := Part1(input)
			g.In1 = override
			g.In2 = override
		}

		gates = append(gates, g)
	}

	state := make(map[string]int)

	for len(gates) > 0 {
		gates = ApplyAll(gates, state)
	}

	return strconv.Itoa(state["a"])
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 46065
	fmt.Println(Part2(aoc.ReadInput())) // 14134
}

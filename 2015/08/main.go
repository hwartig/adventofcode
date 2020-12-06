package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"../../aoc"
)

type Code struct {
	Encoded, Decoded string
}

func (c Code) String() string {
	return fmt.Sprintf("%d : %s\n%d : %s", utf8.RuneCountInString(c.Encoded), c.Encoded, utf8.RuneCountInString(c.Decoded), c.Decoded)
}

func (c Code) Diff() int {
	return utf8.RuneCountInString(c.Encoded) - utf8.RuneCountInString(c.Decoded)
}

func Decode(line string) Code {
	var decoded string

	for i := 0; i < len(line); i++ {
		current := line[i]

		if current == '"' {
			continue
		}

		if current == '\\' {
			next := line[i+1]

			if next == '"' {
				current = '"'
				i += 1
			}

			if next == '\\' {
				i += 1
			}

			if next == 'x' {
				value, _ := strconv.ParseInt(line[i+2:i+4], 16, 64)
				current = byte(value)

				i += 3
			}
		}

		decoded += string(current)
	}

	return Code{line, decoded}
}

func Encode(line string) Code {
	encoded := "\""

	for i := 0; i < len(line); i++ {
		current := line[i]

		if current == '\\' {
			encoded += "\\\\"
		} else if current == '"' {
			encoded += "\\\""
		} else {
			encoded += string(current)
		}
	}

	encoded += "\""

	return Code{encoded, line}
}

func Part1(input string) string {
	result := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		result += Decode(line).Diff()
	}

	return strconv.Itoa(result)
}

func Part2(input string) string {
	result := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		result += Encode(line).Diff()
	}

	return strconv.Itoa(result)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 1350
	fmt.Println(Part2(aoc.ReadInput())) // 2085
}

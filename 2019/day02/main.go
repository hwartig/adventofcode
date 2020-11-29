package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var fileName = "input.txt"

func run(s string, noun, verb int) []int {
	input := splitToNumbers(s)

	// add overrides to original input
	if noun > 0 {
		input[1] = noun
	}
	if verb > 0 {
		input[2] = verb
	}

	pos := 0
	for {
		instr := input[pos]

		if instr == 99 {
			return input
		}

		a := input[input[pos+1]]
		b := input[input[pos+2]]

		if instr == 1 {
			input[input[pos+3]] = a + b
		}

		if instr == 2 {
			input[input[pos+3]] = a * b
		}

		pos = pos + 4
	}
}

func splitToNumbers(s string) (numbers []int) {
	numberStrings := strings.Split(s, ",")
	numbers = make([]int, len(numberStrings))

	for i, n := range numberStrings {
		number, _ := strconv.Atoi(n)
		numbers[i] = number
	}
	return
}

func part1(input string) {
	fmt.Println(run(input, 12, 2)[0])
}

func part2(input string) {
	expectedOutput := 19690720

	for noun := 1; noun < 100; noun++ {
		for verb := 1; verb < 100; verb++ {
			result := run(input, noun, verb)
			if result[0] == expectedOutput {
				fmt.Println(noun*100 + verb)
			}
		}
	}
}

func main() {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("couldn't read file: ", fileName, err)
	}

	//part1(string(content))

	part2(string(content))
}

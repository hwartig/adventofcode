package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var fileName = "input.txt"

func run(s string) []int {
	input := splitToNumbers(s)

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

func main() {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("couldn't read file: ", fileName, err)
	}

	fmt.Println(run(string(content))[0])
}

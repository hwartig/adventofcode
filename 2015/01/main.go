package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func Part1(input string) string {
	result := 0
	for _, c := range input {
		if c == '(' {
			result += 1
		} else if c == ')' {
			result -= 1
		} else {
			//	log.Fatal("found unknown rune", c)
		}
	}
	return strconv.Itoa(result)
}

func Part2(input string) string {
	result := 0
	for i, c := range input {
		if c == '(' {
			result += 1
		} else if c == ')' {
			result -= 1
		}

		if result == -1 {
			return strconv.Itoa(i + 1)
		}
	}
	return strconv.Itoa(-1)
}

func ReadInput() string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func main() {
	fmt.Println(Part2(ReadInput()))
}

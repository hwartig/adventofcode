package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Part1(input string) string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		num1 := Atoi(line)
		for j := i + 1; j < len(lines); j++ {
			num2 := Atoi(lines[j])

			if num1+num2 == 2020 {
				return strconv.Itoa(num1 * num2)
			}
		}
	}
	return ""
}

func Part2(input string) string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		num1 := Atoi(line)

		for j := i + 1; j < len(lines); j++ {
			num2 := Atoi(lines[j])

			for k := j; k < len(lines); k++ {
				num3 := Atoi(lines[k])

				if num1+num2+num3 == 2020 {
					return strconv.Itoa(num1 * num2 * num3)
				}
			}
		}
	}
	return ""
}

func ReadInput() string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(input), "\n")
}

func main() {
	fmt.Println(Part1(ReadInput()))
	fmt.Println(Part2(ReadInput()))
}

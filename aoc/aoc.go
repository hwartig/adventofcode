package aoc

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Max(a []int) int {
	_, max := MinMax(a)
	return max
}

func Min(a []int) int {
	min, _ := MinMax(a)
	return min
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]

	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	return min, max
}

func ReadInput() string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(input), "\n")
}

func Sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

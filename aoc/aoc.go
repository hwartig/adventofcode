package aoc

import (
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func AssertEq(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		log.Println("AssertEq failed")
		log.Printf("%#v != %#v\n", expected, actual)
	}
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

func Multiply(values []int) int {
	product := 1
	for _, value := range values {
		product *= value
	}
	return product
}

type Pair struct {
	Key, Val int
}

type ByKey []Pair

func (s ByKey) Len() int {
	return len(s)
}

func (s ByKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByKey) Less(i, j int) bool {
	return s[i].Val < s[j].Val
}

type ByValue []Pair

func (s ByValue) Len() int {
	return len(s)
}

func (s ByValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByValue) Less(i, j int) bool {
	return s[i].Val < s[j].Val
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

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"../../aoc"
)

func decString(code string) int {
	from := 0
	to := int(math.Pow(2, float64(len(code)))) - 1

	for _, c := range code {
		from, to = decRune(from, to, c)
	}

	return from
}

func decRune(from, to int, c rune) (int, int) {
	mid := from + (to-from)/2

	if c == 'F' || c == 'L' {
		to = mid
	} else if c == 'B' || c == 'R' {
		from = mid + 1
	}

	return from, to
}

func SeatIds(input string) ([]int, int) {
	var ids []int

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		row := decString(line[0:7])
		col := decString(line[7:10])

		ids = append(ids, row*8+col)
	}

	sort.Ints(ids)

	return ids, ids[len(ids)-1]
}

func Part1(input string) string {
	_, maxId := SeatIds(input)

	return strconv.Itoa(maxId)
}

func Part2(input string) string {
	ids, _ := SeatIds(input)

	for i, id := range ids {
		if id+2 == ids[i+1] {
			return strconv.Itoa(id + 1)
		}
	}

	return "not found"
}

func main() {
	//fmt.Println(dec(0, 127, 'F'))
	//fmt.Println(dec(0, 63, 'B'))
	//fmt.Println(dec(32, 63, 'F'))
	//fmt.Println(dec(32, 47, 'B'))
	//fmt.Println(dec(40, 47, 'B'))
	//fmt.Println(dec(44, 47, 'F'))
	//fmt.Println(dec(44, 45, 'F'))
	//fmt.Println(dec(0, 7, 'R'))
	//fmt.Println(dec(4, 7, 'L'))
	//fmt.Println(dec(4, 5, 'R'))
	//fmt.Println(Part1("FBFBBFFRLR"))
	//fmt.Println(Part1("BFFFBBFRRR"))
	//fmt.Println(Part1("FFFBBBFRRR"))
	//fmt.Println(Part1("BBFFBBFRLL"))

	fmt.Println(Part1(aoc.ReadInput())) // 938
	fmt.Println(Part2(aoc.ReadInput())) // 696
}

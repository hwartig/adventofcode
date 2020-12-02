package main

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"../../aoc"
)

func FindHashWithLeadingZeroes(input string, requiredZeroes int) string {
	leadingZeroes := fmt.Sprintf("%0"+strconv.Itoa(requiredZeroes)+"d", 0)

	for i := 0; i <= 10000000; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		if fmt.Sprintf("%x", hash)[0:requiredZeroes] == leadingZeroes {
			return strconv.Itoa(i)
		}
	}

	return "reached limit"
}

func Part1(input string) string {
	return FindHashWithLeadingZeroes(input, 5)
}

func Part2(input string) string {
	return FindHashWithLeadingZeroes(input, 6)
}

func main() {
	//fmt.Println(Part1("abcdef"))  // 609043
	//fmt.Println(Part1("pqrstuv")) // 1048970

	fmt.Println(Part1(aoc.ReadInput())) // 282749
	fmt.Println(Part2(aoc.ReadInput())) // 9962624
}

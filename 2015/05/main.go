package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func HasThreeOrMoreVowels(word string) bool {
	vowels := 0
	for _, c := range word {
		if strings.ContainsRune("aeiou", c) {
			vowels += 1
		}
	}
	return vowels >= 3
}

func HasDoubleLetter(word string) bool {
	prev := ' '
	for i, c := range word {
		if i > 0 && prev == c {
			return true
		}
		prev = c
	}
	return false
}

func HasForbiddenStrings(word string) bool {
	return strings.Contains(word, "ab") ||
		strings.Contains(word, "cd") ||
		strings.Contains(word, "pq") ||
		strings.Contains(word, "xy")
}

func Part1(input string) string {
	result := 0
	lines := strings.Split(input, "\n")
	for _, word := range lines {
		if HasThreeOrMoreVowels(word) && HasDoubleLetter(word) && !HasForbiddenStrings(word) {
			result += 1
		}
	}
	return strconv.Itoa(result)
}

func HasPair(word string) bool {
	prev := ' '
	for i, c := range word {
		if i > 0 && strings.Count(word, string([]rune{prev, c})) >= 2 {
			return true
		}
		prev = c
	}
	return false
}

func HasTriple(word string) bool {
	prev1 := ' '
	prev2 := ' '
	for i, c := range word {
		if i > 1 && prev2 == c {
			return true
		}
		prev2 = prev1
		prev1 = c
	}
	return false
}

func Part2(input string) string {
	result := 0
	lines := strings.Split(input, "\n")
	for _, word := range lines {
		if HasPair(word) && HasTriple(word) {
			result += 1
		}
	}
	return strconv.Itoa(result)
}

func main() {
	//fmt.Println(Part1("ugknbfddgicrmopn"))
	//fmt.Println(Part1("aaa"))
	//fmt.Println(Part1("jchzalrnumimnmhp"))
	//fmt.Println(Part1("haegwjzuvuyypxyu"))
	//fmt.Println(Part1("dvszwmarrgswjxmb"))

	fmt.Println(Part1(aoc.ReadInput())) // 236

	//fmt.Println(Part2("qjhvhtzxzqqjkmpb"))
	//fmt.Println(Part2("xxyxx"))
	//fmt.Println(Part2("uurcxstgmygtbstg"))
	//fmt.Println(Part2("ieodomkazucvgmuy"))

	fmt.Println(Part2(aoc.ReadInput())) // 51
}

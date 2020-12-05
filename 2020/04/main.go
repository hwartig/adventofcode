package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../../aoc"
)

func ValidPart1(p string) bool {
	return strings.Contains(p, "byr") && // Birth Year
		strings.Contains(p, "iyr") && // Issue Year
		strings.Contains(p, "eyr") && // Expiration Year
		strings.Contains(p, "hgt") && // Height
		strings.Contains(p, "hcl") && // Hair Color
		strings.Contains(p, "ecl") && // Eye Color
		strings.Contains(p, "pid") // Passport ID
}

func IsNumberBetween(value string, low, high int) bool {
	num := aoc.Atoi(value)

	return num >= low && num <= high
}

func ValidPart2(p string) bool {
	valid := ValidPart1(p)

	parts := regexp.MustCompile("\\s").Split(p, -1)

	for i := 0; i < len(parts) && valid; i++ {
		kv := strings.Split(parts[i], ":")
		key, value := kv[0], kv[1]

		switch key {
		case "byr":
			valid = IsNumberBetween(value, 1920, 2002)
		case "iyr":
			valid = IsNumberBetween(value, 2010, 2020)
		case "eyr":
			valid = IsNumberBetween(value, 2020, 2030)
		case "hgt":
			if strings.HasSuffix(value, "cm") {
				valid = IsNumberBetween(value[0:len(value)-2], 150, 193)
			} else if strings.HasSuffix(value, "in") {
				valid = IsNumberBetween(value[0:len(value)-2], 59, 76)
			} else {
				valid = false
			}
		case "hcl":
			valid = regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(value)
		case "ecl":
			valid = regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$").MatchString(value)
		case "pid":
			valid = regexp.MustCompile("^\\d{9}$").MatchString(value)
		}
	}

	return valid
}

func Part1(input string) string {
	result := 0

	passports := strings.Split(input, "\n\n")

	for _, p := range passports {

		if ValidPart1(p) {
			result += 1
		}
	}

	return strconv.Itoa(result)
}

func Part2(input string) string {
	result := 0

	passports := strings.Split(input, "\n\n")

	for _, p := range passports {

		if ValidPart2(p) {
			result += 1
		}
	}

	return strconv.Itoa(result)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 242
	fmt.Println(Part2(aoc.ReadInput())) // 186
}

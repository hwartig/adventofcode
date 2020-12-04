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
	//strings.Contains(p, "cid") // Country ID ignored
}

func ValidPart2(p string) bool {
	valid := ValidPart1(p)

	parts := regexp.MustCompile("\\s").Split(p, -1)

	for i := 0; i < len(parts) && valid; i++ {
		kv := strings.Split(parts[i], ":")
		key, value := kv[0], kv[1]

		switch key {
		case "byr":
			year, err := strconv.Atoi(value)
			valid = err == nil && year >= 1920 && year <= 2002
		case "iyr":
			year, err := strconv.Atoi(value)
			valid = err == nil && year >= 2010 && year <= 2020
		case "eyr":
			year, err := strconv.Atoi(value)
			valid = err == nil && year >= 2020 && year <= 2030
		case "hgt":
			if strings.HasSuffix(value, "cm") {
				height, err := strconv.Atoi(value[0 : len(value)-2])
				valid = err == nil && height >= 150 && height <= 193
			} else if strings.HasSuffix(value, "in") {
				height, err := strconv.Atoi(value[0 : len(value)-2])
				valid = err == nil && height >= 59 && height <= 76
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

	fmt.Println(Part2(`eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

eyr:2029 ecl:blu cid:129 byr:2002 iyr:2014 pid:896056539 hcl:#a97842 hgt:190in

eyr:2029 ecl:blu cid:129 byr:2002 iyr:2014 pid:896056539 hcl:#a97842 hgt:190

eyr:2029 ecl:blu cid:129 byr:2003 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#123abz hgt:165cm

eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:123abz hgt:165cm

eyr:2029 ecl:wat cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#123abc hgt:165cm

eyr:2029 ecl:brn cid:129 byr:1989 iyr:2014 pid:0123456789 hcl:#123abc hgt:165cm

hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007`))

	fmt.Println(Part2(`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f

eyr:2029 ecl:brn cid:129 byr:1989 iyr:2014 pid:000000001 hcl:#123abc hgt:165cm

eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#123abc hgt:165cm

eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

eyr:2029 ecl:blu cid:129 byr:2002 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

eyr:2029 ecl:blu cid:129 byr:2002 iyr:2014 pid:896056539 hcl:#a97842 hgt:60in

eyr:2029 ecl:blu cid:129 byr:2002 iyr:2014 pid:896056539 hcl:#a97842 hgt:190cm

hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`))
}

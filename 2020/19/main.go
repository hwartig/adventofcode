package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../../aoc"
)

func Resolve(rules []string, ruleNo int) string {
	rule := rules[ruleNo]
	if rule[0] == '"' && rule[2] == '"' {
		return string(rule[1])
	} else {
		var result []string
		parts := strings.Split(rule, " | ")

		for _, p := range parts {
			nums := strings.Split(p, " ")

			pattern := ""

			for _, n := range nums {
				pattern += Resolve(rules, aoc.Atoi(n))
			}

			result = append(result, pattern)

		}
		return "(" + strings.Join(result, "|") + ")"
	}
}

func ResolveWithLoop(rules []string, ruleNo int) string {
	if ruleNo == 8 {
		return ResolveWithLoop(rules, 42) + "+"
	}

	if ruleNo == 11 {
		first := ResolveWithLoop(rules, 42)
		second := ResolveWithLoop(rules, 31)

		var res []string

		// max line length < 100
		for i := 1; i < 50; i++ {
			res = append(res, fmt.Sprintf("%s{%d}%s{%d}", first, i, second, i))
		}
		return "(" + strings.Join(res, "|") + ")"
	}

	rule := rules[ruleNo]

	if rule[0] == '"' && rule[2] == '"' {
		return string(rule[1])
	} else {
		var result []string
		parts := strings.Split(rule, " | ")

		for _, p := range parts {
			nums := strings.Split(p, " ")

			pattern := ""

			for _, n := range nums {
				pattern += ResolveWithLoop(rules, aoc.Atoi(n))
			}

			result = append(result, pattern)
		}

		return "(" + strings.Join(result, "|") + ")"
	}
}

func Parse(input string) ([]string, []string) {
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")
	rules := make([]string, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rules[aoc.Atoi(parts[0])] = parts[1]
	}

	messages := strings.Split(parts[1], "\n")

	return rules, messages
}

func CountMatches(messages []string, re *regexp.Regexp) (result int) {
	for _, line := range messages {
		if re.MatchString(line) {
			result += 1
		}
	}

	return result
}

func Part1(input string) string {
	rules, messages := Parse(input)

	re := regexp.MustCompile("^" + Resolve(rules, 0) + "$")

	return strconv.Itoa(CountMatches(messages, re))
}

func Part2(input string) string {
	rules, messages := Parse(input)

	re := regexp.MustCompile("^" + ResolveWithLoop(rules, 0) + "$")

	return strconv.Itoa(CountMatches(messages, re))
}

func main() {
	sample := `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

	aoc.AssertEq("2", Part1(sample))

	fmt.Println(Part1(aoc.ReadInput())) // 279
	fmt.Println(Part2(aoc.ReadInput())) // 384
}

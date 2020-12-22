package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

type Validation struct {
	Name                            string
	Low1, High1, Low2, High2, Field int
	Fields                          []int
}

func ParseRules(input string) (result Validations) {
	lines := strings.Split(input, "\n")

	for _, l := range lines {
		parts := strings.Split(l, ": ")
		ranges := strings.Split(parts[1], " or ")

		var nums []int
		for _, r := range ranges {
			ns := strings.Split(r, "-")
			nums = append(nums, aoc.Atoi(ns[0]))
			nums = append(nums, aoc.Atoi(ns[1]))
		}
		result = append(result, &Validation{parts[0], nums[0], nums[1], nums[2], nums[3], -1, []int{}})
	}

	return result
}

func (v Validation) CheckNum(num int) bool {
	return (num >= v.Low1 && num <= v.High1) || (num >= v.Low2 && num <= v.High2)
}

type Validations []*Validation

func (vals Validations) CheckTicket(ticket string) (result int, nums []int) {
	parts := strings.Split(ticket, ",")

	for _, p := range parts {
		nums = append(nums, aoc.Atoi(p))
	}

	for _, num := range nums {
		valid := false
		for _, v := range vals {
			if v.CheckNum(num) {
				valid = true
			}
		}
		if !valid {
			result += num
		}
	}
	return result, nums
}

func (vals Validations) Remove(val int) {
	for _, v := range vals {

		var new []int

		for _, n := range v.Fields {
			if n != val {
				new = append(new, n)
			}
		}
		v.Fields = new
	}
}

func Part1(input string) string {
	blocks := strings.Split(input, "\n\n")

	validations := ParseRules(blocks[0])

	scanErrRate := 0

	lines := strings.Split(blocks[2], "\n")

	for i, ticket := range lines {
		if i == 0 {
			continue
		}
		scanErr, _ := validations.CheckTicket(ticket)
		scanErrRate += scanErr
	}

	return strconv.Itoa(scanErrRate)
}

func Part2(input string) string {
	result := 1

	blocks := strings.Split(input, "\n\n")
	validations := ParseRules(blocks[0])

	lines := strings.Split(blocks[2], "\n")

	var validTickets [][]int

	for i, ticket := range lines {
		if i == 0 {
			continue
		}
		if err, nums := validations.CheckTicket(ticket); err == 0 {
			validTickets = append(validTickets, nums)
		}
	}

	yourTicket := strings.TrimPrefix(blocks[1], "your ticket:\n")
	_, nums := validations.CheckTicket(yourTicket)
	validTickets = append(validTickets, nums)

	for i := 0; i < len(validTickets[0]); i++ {
		for _, v := range validations {
			valid := true

			for j := 0; j < len(validTickets); j++ {
				valid = valid && v.CheckNum(validTickets[j][i])
			}

			if valid {
				v.Fields = append(v.Fields, i)
			}
		}
	}

	// luckily my input doesn't contain ambiguous field mapping
	for i, j := 0, 1; j < len(validations); i = (i + 1) % len(validations) {
		v := validations[i]

		if len(v.Fields) == 1 {
			v.Field = v.Fields[0]
			validations.Remove(v.Field)
			j++
		}
	}

	for _, v := range validations {
		if strings.HasPrefix(v.Name, "departure") {
			result *= nums[v.Field]
		}
	}

	return strconv.Itoa(result)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput())) // 21081
	fmt.Println(Part2(aoc.ReadInput())) // 314360510573
}

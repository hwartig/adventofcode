package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var fileName = "input.txt"

type instruction struct {
	OpCode                 int
	Param1, Param2, Param3 int
	Mode1, Mode2, Mode3    int
}

func parseInstruction(input []int) instruction {
	var p1, p2, p3 int

	if len(input) > 1 {
		p1 = input[1]
	}
	if len(input) > 2 {
		p2 = input[2]
	}
	if len(input) > 3 {
		p3 = input[3]
	}

	modesAndOptCode := fmt.Sprintf("%05d", input[0])
	// mode flags are in decreasing order from left to right
	mode1, _ := strconv.Atoi(string(modesAndOptCode[2]))
	mode2, _ := strconv.Atoi(string(modesAndOptCode[1]))
	mode3, _ := strconv.Atoi(string(modesAndOptCode[0]))
	opCode, _ := strconv.Atoi(string(modesAndOptCode[3:]))

	return instruction{
		OpCode: opCode,
		Param1: p1,
		Param2: p2,
		Param3: p3,
		Mode1:  mode1,
		Mode2:  mode2,
		Mode3:  mode3,
	}
}

func resolveParams(instr *instruction, input []int, count int) {
	if count > 0 && instr.Mode1 == 0 { // not really needed. why would you call the method if count == 0
		instr.Param1 = input[instr.Param1]
	}
	if count > 1 && instr.Mode2 == 0 {
		instr.Param2 = input[instr.Param2]
	}
	if count > 2 && instr.Mode3 == 0 {
		instr.Param3 = input[instr.Param3]
	}
}

func run(s string, noun, verb int) []int {
	input := splitToNumbers(s)

	// add some padding to the end as we're parsing instrs in batches of 4
	input = append(input, 0, 0, 0)

	// add overrides to original input
	if noun > 0 {
		input[1] = noun
	}
	if verb > 0 {
		input[2] = verb
	}

	pos := 0
	for {
		instr := parseInstruction(input[pos : pos+4])
		// log.Println("instruction: ", instr)

		switch instr.OpCode {
		case 99: // halt
			return input[:len(input)-3] // return without extra padding
		case 1: // addition
			resolveParams(&instr, input, 2)
			input[instr.Param3] = instr.Param1 + instr.Param2
			pos = pos + 4
		case 2: // multiplication
			resolveParams(&instr, input, 2)
			input[instr.Param3] = instr.Param1 * instr.Param2
			pos = pos + 4
		case 3: // input
			fmt.Scan(&input[instr.Param1])

			pos = pos + 2
		case 4: // output
			resolveParams(&instr, input, 1)
			fmt.Println(instr.Param1)

			pos = pos + 2
		default:
			log.Fatalf("Unknown opcode %v in instruction %v\n", instr.OpCode, instr)
		}
	}
}

func splitToNumbers(s string) (numbers []int) {
	numberStrings := strings.Split(s, ",")
	numbers = make([]int, len(numberStrings))

	for i, n := range numberStrings {
		number, _ := strconv.Atoi(n)
		numbers[i] = number
	}
	return
}

func main() {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("couldn't read file: ", fileName, err)
	}

	// run("3,0,4,0,99", 0, 0)
	run(string(content), 0, 0)
}

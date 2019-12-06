package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName = "input.txt"

type instruction struct {
	OpCode                 int
	Param1, Param2, Param3 int
	Mode1, Mode2, Mode3    int
}

func parseInstruction(mem []int) instruction {
	var p1, p2, p3 int

	if len(mem) > 1 {
		p1 = mem[1]
	}
	if len(mem) > 2 {
		p2 = mem[2]
	}
	if len(mem) > 3 {
		p3 = mem[3]
	}

	modesAndOptCode := fmt.Sprintf("%05d", mem[0])
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

func resolveParams(instr *instruction, mem []int, count int) {
	if count > 0 && instr.Mode1 == 0 { // not really needed. why would you call the method if count == 0
		instr.Param1 = mem[instr.Param1]
	}
	if count > 1 && instr.Mode2 == 0 {
		instr.Param2 = mem[instr.Param2]
	}
	if count > 2 && instr.Mode3 == 0 {
		instr.Param3 = mem[instr.Param3]
	}
}

func run(s string, in io.Reader, out io.Writer) []int {
	mem := splitToNumbers(s)

	// add some padding to the end as we're parsing instrs in batches of 4
	mem = append(mem, 0, 0, 0)

	pos := 0
	for {
		instr := parseInstruction(mem[pos : pos+4])

		switch instr.OpCode {
		case 99: // halt
			return mem[:len(mem)-3] // return without extra padding
		case 1: // addition
			resolveParams(&instr, mem, 2)
			mem[instr.Param3] = instr.Param1 + instr.Param2
			pos = pos + 4
		case 2: // multiplication
			resolveParams(&instr, mem, 2)
			mem[instr.Param3] = instr.Param1 * instr.Param2
			pos = pos + 4
		case 3: // input
			s := bufio.NewScanner(in)
			if s.Scan() {
				n, err := strconv.Atoi(s.Text())
				if err != nil {
					log.Fatal("Unable to read input, ", err)
				}
				mem[instr.Param1] = n
			}

			pos = pos + 2
		case 4: // output
			resolveParams(&instr, mem, 1)

			_, err := io.WriteString(out, strconv.Itoa(instr.Param1))

			if err != nil {
				log.Fatal("Unable to write to output, ", err)
			}

			pos = pos + 2
		case 5: // jump if true
			resolveParams(&instr, mem, 2)
			if instr.Param1 != 0 {
				pos = instr.Param2
			} else {
				pos = pos + 3
			}
		case 6: // jump if false
			resolveParams(&instr, mem, 2)
			if instr.Param1 == 0 {
				pos = instr.Param2
			} else {
				pos = pos + 3
			}
		case 7: // less than
			resolveParams(&instr, mem, 2)
			if instr.Param1 < instr.Param2 {
				mem[instr.Param3] = 1
			} else {
				mem[instr.Param3] = 0
			}
			pos = pos + 4
		case 8: // equals
			resolveParams(&instr, mem, 2)
			if instr.Param1 == instr.Param2 {
				mem[instr.Param3] = 1
			} else {
				mem[instr.Param3] = 0
			}
			pos = pos + 4
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

	run("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", os.Stdin, os.Stdout)
	run(string(content), os.Stdin, os.Stdout)
}

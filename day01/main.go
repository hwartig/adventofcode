package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func CalcFuel(mass int) int {
	return int(math.Floor(float64(mass/3)) - 2)
}

var fileName = "input.txt"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("could not open file: ", fileName, err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringMass := scanner.Text()
		mass, err := strconv.Atoi(stringMass)

		if err != nil {
			log.Fatalf("could not convert %v to integer", stringMass)
		}

		total = total + CalcFuel(mass)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("reading standard input:", err)
	}
	fmt.Println(total)
}

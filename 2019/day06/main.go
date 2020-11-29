package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Planet struct {
	Name   string
	Orbits *Planet
}

func (plt *Planet) CountOrbits() (count int) {

	for p := plt; p.Orbits != nil; p = p.Orbits {
		count++
	}

	return
}

func parseInput(r io.Reader) map[string]*Planet {
	s := bufio.NewScanner(r)
	planets := make(map[string]*Planet)

	for i := 0; s.Scan(); i++ {
		ps := strings.Split(s.Text(), ")")

		if len(ps) != 2 {
			log.Fatalf("cannot read line %d : %v", i, s.Text())
		}

		p1, ok := planets[ps[0]]
		if !ok {
			p1 = &Planet{Name: ps[0]}
			planets[ps[0]] = p1
		}

		p2, ok := planets[ps[1]]
		if ok {
			p2.Orbits = p1
		} else {
			p2 = &Planet{ps[1], p1}
			planets[ps[1]] = p2
		}
	}
	return planets
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("couldn't open file ", err)
	}

	planets := parseInput(f)

	total := 0

	for _, p := range planets {
		total += p.CountOrbits()
	}

	fmt.Println("part1: ", total)
}

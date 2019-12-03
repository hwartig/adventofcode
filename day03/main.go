package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Step struct {
	Direction string
	Distance  int
}

type Point struct {
	X int
	Y int
}

func (a Point) Compare (b Point) int {
  if a.ManhattanDistance() < b.ManhattanDistance() {
    return -1
  } 

  if a.ManhattanDistance() > b.ManhattanDistance() {
    return 1
  }

  if a.X < b.X {
    return -1
  }

  if a.X > b.X {
    return 1
  }

  if a.Y < b.Y {
    return -1
  }
  
  if a.Y > b.Y {
    return 1
  }

  return 0
}

type Points []Point

func (p Points) Len() int {
  return len(p)
}
func (p Points) Swap(i, j int) {
  p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
  return p[i].Compare(p[j]) < 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func(p Point) ManhattanDistance() int {
  return abs(p.X) + abs(p.Y)
}

type Wire struct {
  Steps []Step
}

func (w Wire) Points() (points Points) {
	x, y := 0, 0 // starting position

	for _, step := range w.Steps {
		for i := 0; i < step.Distance; i += 1 {
			switch step.Direction {
			case "U":
				y = y + 1
			case "D":
				y = y - 1
			case "R":
				x = x + 1
			case "L":
				x = x - 1
			}

			points = append(points, Point{x, y})
		}
	}
  sort.Sort(points)

	return
}

func parseWire(input string) (wire Wire) {
	parts := strings.Split(input, ",")

	for _, part := range parts {
		direction := string(part[0])
		number, _ := strconv.Atoi(part[1:])
    wire.Steps = append(wire.Steps, Step{direction, number})
	}

	return
}

func (a Wire) FindOverlaps(b Wire) (points Points) {
  ap := a.Points()
  bp := b.Points()

  for i, j := 0, 0; i < len(ap) && j < len(bp); {
    //fmt.Printf("%d: %v < %d: %v\n", i, ap[i], j, bp[j])

    switch(ap[i].Compare(bp[j])) {
    case -1:
      i = i + 1
    case 0:
      points = append(points, ap[i])
      i = i + 1
      j = j + 1
    case 1:
      j = j + 1
    } 
  }

  return 
}

func main() {
  file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("couldn't open file, ", err)
	}
  defer file.Close()

  s := bufio.NewScanner(file)
  s.Scan()
  wire1 := parseWire(s.Text())
  s.Scan()
  wire2 := parseWire(s.Text())

  fmt.Println(wire1.FindOverlaps(wire2)[0].ManhattanDistance())
}

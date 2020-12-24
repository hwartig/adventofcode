package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

type AxialVector struct {
	q, r int
}

func (v AxialVector) Parse(in string) AxialVector {
	if len(in) == 0 {
		return v
	}

	if in[0] == 'e' {
		return AxialVector{v.q + 1, v.r}.Parse(in[1:])
	} else if in[0] == 'w' {
		return AxialVector{v.q - 1, v.r}.Parse(in[1:])
	} else {
		switch in[0:2] {
		case "se":
			return AxialVector{v.q, v.r + 1}.Parse(in[2:])
		case "sw":
			return AxialVector{v.q - 1, v.r + 1}.Parse(in[2:])
		case "ne":
			return AxialVector{v.q + 1, v.r - 1}.Parse(in[2:])
		case "nw":
			return AxialVector{v.q, v.r - 1}.Parse(in[2:])
		}
	}

	return v
}

type Floor map[AxialVector]bool

func NewFloor(input string) Floor {
	floor := make(map[AxialVector]bool)

	for _, line := range strings.Split(input, "\n") {
		v := AxialVector{}.Parse(line)
		floor[v] = !floor[v]
	}

	return floor
}

func (f Floor) CountBlackTiles() (result int) {
	for _, v := range f {
		if v {
			result += 1
		}
	}
	return
}

func (f Floor) MinMaxCoords() (minq, maxq, minr, maxr int) {
	for k, _ := range f {
		if k.q < minq {
			minq = k.q
		}
		if k.q > maxq {
			maxq = k.q
		}
		if k.r < minr {
			minr = k.r
		}
		if k.r > maxr {
			maxr = k.r
		}
	}

	return
}

func (f Floor) Copy() Floor {
	copy := make(Floor)

	for k, v := range f {
		copy[k] = v
	}

	return copy
}

func (f Floor) Step() Floor {
	next := f.Copy()

	minq, maxq, minr, maxr := f.MinMaxCoords()

	for q := minq - 1; q <= maxq+1; q++ {
		//if (q-minq)%2 == 0 {
		//fmt.Print(" ")
		//}
		for r := minr - 1; r <= maxr+1; r++ {
			v := AxialVector{q, r}
			//fmt.Println(v, f[v], f.CountBlackNeighborsTiles(v))
			//fmt.Print(f.CountBlackNeighborsTiles(v), " ")
			//if f[v] {
			//fmt.Print("X ")
			//} else {
			//fmt.Print("  ")
			//}

			if f[v] {
				if f.CountBlackNeighborsTiles(v) == 0 || f.CountBlackNeighborsTiles(v) > 2 {
					delete(next, v)
				}
			} else {
				if f.CountBlackNeighborsTiles(v) == 2 {
					next[v] = true
				}
			}
		}
		//fmt.Println()
	}
	return next
}

func (f Floor) CountBlackNeighborsTiles(v AxialVector) (result int) {
	if f[AxialVector{v.q + 1, v.r}] {
		result += 1
	}
	if f[AxialVector{v.q + 1, v.r - 1}] {
		result += 1
	}
	if f[AxialVector{v.q, v.r - 1}] {
		result += 1
	}
	if f[AxialVector{v.q - 1, v.r}] {
		result += 1
	}
	if f[AxialVector{v.q - 1, v.r + 1}] {
		result += 1
	}
	if f[AxialVector{v.q, v.r + 1}] {
		result += 1
	}

	return result
}

func Part1(input string) string {
	return strconv.Itoa(NewFloor(input).CountBlackTiles())
}

func Part2(input string) string {
	floor := NewFloor(input)

	for i := 1; i <= 100; i++ {
		floor = floor.Step()
		//fmt.Println("Day", i, ":", floor.CountBlackTiles())
	}

	return strconv.Itoa(floor.CountBlackTiles())
}

func main() {
	sample := `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

	fmt.Println(Part1(sample))
	fmt.Println(Part2(sample))

	fmt.Println(Part1(aoc.ReadInput()))
	fmt.Println(Part2(aoc.ReadInput()))
}

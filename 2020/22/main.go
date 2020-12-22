package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"../../aoc"
)

type Player struct {
	Deck []int
}

func (p *Player) Copy(len int) (out Player) {
	out.Deck = append(out.Deck, p.Deck[:len]...)
	return out
}

func (p *Player) Unshift() (val int) {
	val = p.Deck[0]
	p.Deck = p.Deck[1:]

	return val
}

func (p *Player) Push(nums ...int) {
	p.Deck = append(p.Deck, nums...)
}

func (p *Player) HasCards() bool {
	return len(p.Deck) > 0
}

func (p *Player) Score() (res int) {
	for i := len(p.Deck) - 1; i >= 0; i-- {
		res += (len(p.Deck) - i) * p.Deck[i]
	}

	return res
}

func (p *Player) Key() (key string) {
	for _, v := range p.Deck {
		key += strconv.Itoa(v) + ","
	}
	return key
}

func ParsePlayer(input string) (p Player) {
	for _, line := range strings.Split(input, "\n")[1:] {
		p.Push(aoc.Atoi(line))
	}

	return p
}

func Part1(input string) string {
	inputs := strings.Split(input, "\n\n")
	p1 := ParsePlayer(inputs[0])
	p2 := ParsePlayer(inputs[1])

	for p1.HasCards() && p2.HasCards() {
		c1, c2 := p1.Unshift(), p2.Unshift()

		if c1 > c2 {
			p1.Push(c1, c2)
		} else {
			p2.Push(c2, c1)
		}
	}

	if p1.HasCards() {
		logger.Println("Player 1 won:", p1)
		return strconv.Itoa(p1.Score())
	} else {
		logger.Println("Player 2 won:", p2)
		return strconv.Itoa(p2.Score())
	}
}

var game = 0

func Combat(p1, p2 Player) (bool, Player) {
	game += 1
	var playedGames = make(map[string]bool)
	logger.Printf("\n=== Game %d ===\n", game)

	for i := 1; ; i++ {
		key := p1.Key()
		if _, ok := playedGames[key]; ok {
			logger.Println("saw key", p1.Key(), "already")
			return true, p1
		} else {
			playedGames[key] = true
		}

		logger.Printf("\n-- Round %d (Game %d) --\n", i, game)
		logger.Println("Player 1's deck:", p1.Deck)
		logger.Println("Player 2's deck:", p2.Deck)

		c1, c2 := p1.Unshift(), p2.Unshift()
		logger.Println("Player 1 plays:", c1)
		logger.Println("Player 2 plays:", c2)

		if c1 > len(p1.Deck) || c2 > len(p2.Deck) {
			if c1 > c2 {
				logger.Printf("Player 1 wins Round %d of Game %d --\n", i, game)
				p1.Push(c1, c2)
			} else {
				logger.Printf("Player 2 wins Round %d of Game %d --\n", i, game)
				p2.Push(c2, c1)
			}
		} else {
			logger.Println("Playing a sub-game to determine the winner...")
			result, _ := Combat(p1.Copy(c1), p2.Copy(c2))
			if result {
				p1.Push(c1, c2)
				logger.Printf("Player 1 wins Round %d of Game %d --\n", i, game)
			} else {
				p2.Push(c2, c1)
				logger.Printf("Player 2 wins Round %d of Game %d --\n", i, game)
			}
		}

		if !p1.HasCards() {
			logger.Println("The winner of game", game, "is player 2")
			return false, p2
		} else if !p2.HasCards() {
			logger.Println("The winner of game", game, "is player 1")
			return true, p1
		}
	}
}

func Part2(input string) string {

	inputs := strings.Split(input, "\n\n")
	p1 := ParsePlayer(inputs[0])
	p2 := ParsePlayer(inputs[1])

	_, winner := Combat(p1, p2)
	logger.Println("== Post-game results ==")
	logger.Println("Winners deck:", winner.Deck)

	return strconv.Itoa(winner.Score())
}

//var logger = log.New(os.Stderr, "", 0)
var logger = log.New(ioutil.Discard, "", 0)

func main() {
	sample := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

	logger.Println(Part2(sample))
	fmt.Println(Part1(aoc.ReadInput())) // 33925

	logger.Println(Part2(sample))
	fmt.Println(Part2(aoc.ReadInput())) // 33441
}

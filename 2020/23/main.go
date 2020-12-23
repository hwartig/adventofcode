package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

var index []*Cup = make([]*Cup, 1000001)

type Cup struct {
	prev, next *Cup
	val        int
}

func (c *Cup) Add(i int) *Cup {
	new := &Cup{val: i}

	if c == nil {
		new.prev = new
		new.next = new
	} else {
		new.prev = c
		new.next = c.next
		c.next.prev = new
		c.next = new
	}

	index[i] = new

	return new
}

func (c *Cup) AsSlice() (res []int) {
	res = append(res, c.val)

	for cur := c.next; cur != c; cur = cur.next {
		res = append(res, cur.val)
	}

	return res
}

func (c *Cup) String() string {
	res := strconv.Itoa(c.val)

	for cur := c.next; cur != c; cur = cur.next {
		res += ", " + strconv.Itoa(cur.val)
	}

	return res
}

func (c *Cup) AddThree(cups *Cup) {
	cups.next.next.next = c.next
	cups.prev = c

	c.next.prev = cups.next.next
	c.next = cups
}

func (c *Cup) PickUpThree() *Cup {
	start := c.next
	end := start.next.next

	start.prev.next = end.next
	end.next.prev = start.prev

	start.prev = end
	end.next = start

	return start
}

func (c *Cup) Step(min, max int) *Cup {
	cups := c.PickUpThree()

	label := c.val

	for label == c.val || label == cups.val || label == cups.next.val || label == cups.next.next.val {
		label -= 1

		if label < min {
			label = max
		}
	}

	index[label].AddThree(cups)

	return c.next
}

func Part1(input string) string {
	var head *Cup

	for _, n := range strings.Split(input, "") {
		head = head.Add(aoc.Atoi(n))
	}

	head = head.next

	min, max := aoc.MinMax(head.AsSlice())

	for i := 0; i < 100; i++ {
		head = head.Step(min, max)
	}

	return strings.ReplaceAll(index[1].String(), ", ", "")[1:]
}

func Part2(input string) string {
	var head *Cup

	for _, n := range strings.Split(input, "") {
		head = head.Add(aoc.Atoi(n))
	}

	min, max := aoc.MinMax(head.AsSlice())

	for i := max + 1; i <= 1000000; i++ {
		head = head.Add(i)
	}

	head = head.next

	for i := 0; i < 10000000; i++ {
		head = head.Step(min, 1000000)
	}

	head = index[1]

	return strconv.Itoa(head.next.val * head.next.next.val)
}

func main() {
	fmt.Println(Part1(aoc.ReadInput()))
	fmt.Println(Part2(aoc.ReadInput()))
}

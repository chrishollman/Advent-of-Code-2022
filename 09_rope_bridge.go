package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X    int
	Y    int
	Seen map[string]bool
	Head *Point
	Tail *Point
}

type Direction string

const (
	Direction_UP    Direction = "U"
	Direction_DOWN  Direction = "D"
	Direction_LEFT  Direction = "L"
	Direction_RIGHT Direction = "R"
)

func (p *Point) GetStringCoordinates() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}

func (p *Point) Move(d Direction, amount int) {
	switch d {
	case Direction_UP:
		for i := 1; i <= amount; i++ {
			p.Y += 1
			if p.Tail != nil {
				p.updateTailRecursive()
			}
		}
	case Direction_DOWN:
		for i := 1; i <= amount; i++ {
			p.Y -= 1
			if p.Tail != nil {
				p.updateTailRecursive()
			}
		}
	case Direction_LEFT:
		for i := 1; i <= amount; i++ {
			p.X -= 1
			if p.Tail != nil {
				p.updateTailRecursive()
			}
		}
	case Direction_RIGHT:
		for i := 1; i <= amount; i++ {
			p.X += 1
			if p.Tail != nil {
				p.updateTailRecursive()
			}
		}
	}
}

func (p *Point) updateTailRecursive() {
	if p.Tail == nil {
		p.Seen[p.GetStringCoordinates()] = true
		return
	}

	p = p.Tail
	distX := p.Head.X - p.X
	distY := p.Head.Y - p.Y
	absX := abs(distX)
	absY := abs(distY)

	if absX <= 1 && absY <= 1 {
		// Parent within 1 unit of child, no movement necessary and early exit
		return
	} else if (absX == 2 && absY == 0) || (absX == 0 && absY == 2) {
		// Parent within 2 units (absX + absY <= 2) of child in one plane of movement (X or Y, not both)
		switch {
		case distX < 0:
			p.X -= 1
		case distX > 0:
			p.X += 1
		case distY < 0:
			p.Y -= 1
		case distY > 0:
			p.Y += 1
		}
	} else if absX+absY >= 3 {
		// Parent greater 2 units (absX + absY >= 3) of child in more than one plane of movement (X and Y)
		switch {
		case distX < 0:
			p.X -= 1
		case distX > 0:
			p.X += 1
		}
		switch {
		case distY < 0:
			p.Y -= 1
		case distY > 0:
			p.Y += 1
		}
	}

	p.updateTailRecursive()
}

func NewHead() *Point {
	return &Point{
		X:    0,
		Y:    0,
		Seen: nil,
		Head: nil,
		Tail: nil,
	}
}

func NewTail(p *Point) *Point {
	p.Tail = NewHead()
	p.Tail.Head = p
	return p.Tail
}

func dayNineChallengeOne(input []string) int {

	head := NewHead()
	tail := NewTail(head)
	tail.Seen = make(map[string]bool, len(input)*2)
	tail.Seen["0,0"] = true

	for _, line := range input {
		instruction := strings.Fields(line)
		direction := Direction(instruction[0])
		amount, _ := strconv.Atoi(instruction[1])
		head.Move(direction, amount)
	}

	return len(tail.Seen)
}

func dayNineChallengeTwo(input []string) int {
	head := NewHead()
	tail := head
	for i := 1; i < 10; i++ {
		tail = NewTail(tail)
	}
	tail.Seen = make(map[string]bool, len(input)*2)
	tail.Seen["0,0"] = true

	for _, line := range input {
		instruction := strings.Fields(line)
		direction := Direction(instruction[0])
		amount, _ := strconv.Atoi(instruction[1])
		head.Move(direction, amount)
	}

	return len(tail.Seen)
}

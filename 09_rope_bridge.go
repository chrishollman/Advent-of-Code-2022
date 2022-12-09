package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X    int
	Y    int
	Num  int
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

func NewPoint() *Point {
	return &Point{
		X:    0,
		Y:    0,
		Num:  1,
		Seen: map[string]bool{"0,0": true},
		Head: nil,
		Tail: nil,
	}
}

func AddNewTail(p *Point) *Point {
	p.Tail = NewPoint()
	p.Tail.Num = p.Num + 1
	p.Tail.Head = p
	return p.Tail
}

func (p *Point) Move(d Direction, amount int) {
	switch d {
	case Direction_UP:
		p.up(amount)
	case Direction_DOWN:
		p.down(amount)
	case Direction_LEFT:
		p.left(amount)
	case Direction_RIGHT:
		p.right(amount)
	}
}

func (p *Point) up(amount int) {
	for i := 1; i <= amount; i++ {
		p.Y += 1
		if p.Tail != nil {
			p.updateTailRecursive()
		}
	}
}

func (p *Point) down(amount int) {
	for i := 1; i <= amount; i++ {
		p.Y -= 1
		if p.Tail != nil {
			p.updateTailRecursive()
		}
	}
}

func (p *Point) left(amount int) {
	for i := 1; i <= amount; i++ {
		p.X -= 1
		if p.Tail != nil {
			p.updateTailRecursive()
		}
	}
}

func (p *Point) right(amount int) {
	for i := 1; i <= amount; i++ {
		p.X += 1
		if p.Tail != nil {
			p.updateTailRecursive()
		}
	}
}

func (p *Point) getStringCoordinates() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}

func (p *Point) updateTailRecursive() {
	if p.Tail == nil {
		return
	}

	p = p.Tail
	distX := p.Head.X - p.X
	distY := p.Head.Y - p.Y
	absX := abs(distX)
	absY := abs(distY)

	// Parent within 1 unit of child
	if absX <= 1 && absY <= 1 {
		return
		// Parent within 2 units (absX + absY <= 2) of child in one plane of movement (X or Y, not both)
	} else if (absX == 2 && absY == 0) || (absX == 0 && absY == 2) {
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
		// Parent greater 2 units (absX + absY >= 3) of child in more than one plane of movement (X and Y)
	} else if absX+absY >= 3 {
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
	p.Seen[p.getStringCoordinates()] = true
	p.updateTailRecursive()
}

func dayNineChallengeOne(input []string) int {

	head := NewPoint()
	tail := AddNewTail(head)

	for _, line := range input {
		instruction := strings.Fields(line)
		direction := Direction(instruction[0])
		amount, _ := strconv.Atoi(instruction[1])
		head.Move(direction, amount)
	}

	return len(tail.Seen)
}

func dayNineChallengeTwo(input []string) int {
	head := NewPoint()
	tail := head
	for i := 1; i < 10; i++ {
		tail = AddNewTail(tail)
	}

	for _, line := range input {
		instruction := strings.Fields(line)
		direction := Direction(instruction[0])
		amount, _ := strconv.Atoi(instruction[1])
		head.Move(direction, amount)
	}

	return len(tail.Seen)
}

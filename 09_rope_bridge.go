package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

type Leader struct {
	X           int
	Y           int
	Subordinate *Follower
}

type Follower struct {
	X    int
	Y    int
	Seen map[string]bool
}

type Direction string

const (
	Direction_UP    Direction = "U"
	Direction_DOWN  Direction = "D"
	Direction_LEFT  Direction = "L"
	Direction_RIGHT Direction = "R"
)

func NewLeader() *Leader {
	return &Leader{
		X: 0,
		Y: 0,
		Subordinate: &Follower{
			X:    0,
			Y:    0,
			Seen: map[string]bool{"0,0": true},
		},
	}
}

func (l *Leader) Move(d Direction, amount int) {
	switch d {
	case Direction_UP:
		l.up(amount)
	case Direction_DOWN:
		l.down(amount)
	case Direction_LEFT:
		l.left(amount)
	case Direction_RIGHT:
		l.right(amount)
	}
}

func (l *Leader) up(amount int) {
	for i := 1; i <= amount; i++ {
		l.updateSubordinate(l.X, l.Y+1)
		l.Y += 1
	}
}

func (l *Leader) down(amount int) {
	for i := 1; i <= amount; i++ {
		l.updateSubordinate(l.X, l.Y-1)
		l.Y -= 1
	}
}

func (l *Leader) left(amount int) {
	for i := 1; i <= amount; i++ {
		l.updateSubordinate(l.X-1, l.Y)
		l.X -= 1
	}
}

func (l *Leader) right(amount int) {
	for i := 1; i <= amount; i++ {
		l.updateSubordinate(l.X+1, l.Y)
		l.X += 1
	}
}

func (l *Leader) updateSubordinate(newX int, newY int) {

	distX := abs(newX - l.Subordinate.X)
	distY := abs(newY - l.Subordinate.Y)

	if !(distX+distY <= 1) && !(distX == 1 && distY == 1) {
		l.Subordinate.X = l.X
		l.Subordinate.Y = l.Y
		l.Subordinate.Seen[fmt.Sprintf("%v,%v", l.X, l.Y)] = true
	}
}

func dayNineChallengeOne(input []string) int {

	leader := NewLeader()

	for _, line := range input {
		instruction := strings.Fields(line)
		direction := Direction(instruction[0])
		amount, _ := strconv.Atoi(instruction[1])
		leader.Move(direction, amount)
	}

	return len(leader.Subordinate.Seen)
}

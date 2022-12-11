package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	ctr         int
	regX        int
	historicalX []int
	image       [6][40]rune
	postCycle   Queue[instruction]
}

type instruction struct {
	command command
	value   int
}

type command string

const (
	noop command = "noop"
	addx command = "addx"
)

func NewCPU() *cpu {
	return &cpu{
		ctr:         1,
		regX:        1,
		historicalX: []int{1},
		image:       [6][40]rune{},
		postCycle:   Queue[instruction]{},
	}
}

func (c *cpu) Command(com command, v int) {
	switch com {
	case addx:
		c.postCycle.Enqueue(instruction{
			command: noop,
			value:   -1,
		})
		c.postCycle.Enqueue(instruction{
			command: addx,
			value:   v,
		})
	case noop:
		c.postCycle.Enqueue(instruction{
			command: noop,
			value:   -1,
		})
	}
}

func (c *cpu) Cycle() {

	// Draw image pixel
	line := (c.ctr - 1) / 40
	pos := (c.ctr - 1) % 40
	if abs(c.regX-pos) <= 1 {
		c.image[line][pos] = '#'
	} else {
		c.image[line][pos] = '.'
	}

	// Increment cycle counter
	c.ctr++

	// Run postcycle functions
	if !c.postCycle.IsEmpty() {
		instr, _ := c.postCycle.Dequeue()
		c.Exec(instr.command, instr.value)
	}

	// Update register historical value
	c.historicalX = append(c.historicalX, c.regX)
}

func (c *cpu) Exec(com command, v int) {
	switch com {
	case addx:
		c.regX += v
		return
	default:
		return
	}
}

func (c *cpu) RunCycleLoop() {
	for !c.postCycle.IsEmpty() {
		c.Cycle()
	}
}

func runCPU(input []string) *cpu {
	cpu := NewCPU()
	for _, line := range input {
		spl := strings.Fields(line)
		switch {
		case spl[0] == "addx":
			val, _ := strconv.Atoi(spl[1])
			cpu.Command(addx, val)
		case spl[0] == "noop":
			cpu.Command(noop, -1)
		}
	}

	cpu.RunCycleLoop()
	return cpu
}

func dayTenChallengeOne(input []string) int {
	cpu := runCPU(input)

	total := 0
	for i := 20; i <= 220; i = i + 40 {
		total += cpu.historicalX[i-1] * i
	}

	return total
}

func dayTenChallengeTwo(input []string) int {
	cpu := runCPU(input)

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			fmt.Printf("%v", string(cpu.image[i][j]))
		}
		fmt.Printf("\n")
	}

	return -1
}

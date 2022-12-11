package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

type cpu struct {
	cycle           int
	register        int
	registerHistory []int
	image           [6][40]rune
	postCycle       Queue[instruction]
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
		cycle:           1,
		register:        1,
		registerHistory: []int{1},
		image:           [6][40]rune{},
		postCycle:       Queue[instruction]{},
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

func (c *cpu) Exec(com command, v int) {
	switch com {
	case addx:
		c.register += v
		return
	default:
		return
	}
}

func (c *cpu) RunCycle() {
	c.UpdateImageBuffer()
	c.cycle++

	// Run postcycle commands
	if !c.postCycle.IsEmpty() {
		instr, _ := c.postCycle.Dequeue()
		c.Exec(instr.command, instr.value)
	}

	// Update register history
	c.registerHistory = append(c.registerHistory, c.register)
}

func (c *cpu) RunCycleLoop() {
	for !c.postCycle.IsEmpty() {
		c.RunCycle()
	}
}

func (c *cpu) UpdateImageBuffer() {
	line := (c.cycle - 1) / 40
	pos := (c.cycle - 1) % 40
	if abs(c.register-pos) <= 1 {
		c.image[line][pos] = '#'
	} else {
		c.image[line][pos] = '.'
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
		total += cpu.registerHistory[i-1] * i
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

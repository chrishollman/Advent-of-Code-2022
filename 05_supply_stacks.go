package adventofcode2022

import (
	"strconv"
	"strings"
)

func dayFiveChallengeOne(input []string) string {

	inputStacks, inputMovements := splitInstructions(input)
	stacks := getSupplyStacks(inputStacks)
	moves := parseStackMovements(inputMovements)

	for _, move := range moves {
		numCrates := move[0]
		fromStack := move[1] - 1 // Account for 0 index
		toStack := move[2] - 1   // Account for 0 index

		for i := numCrates; i > 0; i-- {
			if val, ok := stacks[fromStack].Pop(); ok {
				stacks[toStack].Push(val)
			}
		}
	}

	var res string
	for _, stack := range stacks {
		if val, ok := stack.Pop(); ok {
			res = res + val
		}
	}

	return res
}

func dayFiveChallengeTwo(input []string) string {

	inputStacks, inputMovements := splitInstructions(input)
	stacks := getSupplyStacks(inputStacks)
	moves := parseStackMovements(inputMovements)

	for _, move := range moves {
		numCrates := move[0]
		fromStack := move[1] - 1 // Account for 0 index
		toStack := move[2] - 1   // Account for 0 index

		var tmpStack Stack
		for i := numCrates; i > 0; i-- {
			if val, ok := stacks[fromStack].Pop(); ok {
				tmpStack.Push(val)
			}
		}

		for j := numCrates; j > 0; j-- {
			if val, ok := tmpStack.Pop(); ok {
				stacks[toStack].Push(val)
			}
		}
	}

	var res string
	for _, stack := range stacks {
		if val, ok := stack.Pop(); ok {
			res = res + val
		}
	}

	return res
}

func getSupplyStacks(input []string) []Stack {
	maxIdx := len(input[0]) - 2
	stacks := make([]Stack, (len(input[0])/4 + 1))

	for i := len(input) - 1; i >= 0; i-- {
		for j := 1; j <= maxIdx; j += 4 {
			val := string(input[i][j])
			if val != " " {
				stacks[j/4].Push(val)
			}
		}
	}

	return stacks
}

func parseStackMovements(input []string) [][]int {
	movements := make([][]int, len(input))
	for idx, plaintext := range input {
		strSplit := strings.Split(plaintext, " ")
		first, _ := strconv.Atoi(strSplit[1])
		second, _ := strconv.Atoi(strSplit[3])
		third, _ := strconv.Atoi(strSplit[5])
		movements[idx] = []int{first, second, third}
	}

	return movements
}

func splitInstructions(input []string) ([]string, []string) {

	var stacks []string
	var movements []string

	for idx, line := range input {
		if strings.HasPrefix(line, " 1") {
			stacks = input[:idx]
			movements = input[idx+2:]
			break
		}
	}

	return stacks, movements
}

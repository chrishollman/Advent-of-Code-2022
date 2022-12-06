package adventofcode2022

import (
	"strings"
	"unicode"
)

func dayThreeChallengeOne(input []string) int {
	var res int
	for _, rucksack := range input {
		midpoint := len(rucksack) / 2
		left, right := strings.Split(rucksack[:midpoint], ""), strings.Split(rucksack[midpoint:], "")
		intersection := SliceIntersection(left, right)
		res += letterToValue(intersection)
	}
	return res
}

func dayThreeChallengeTwo(input []string) int {
	var res int
	for i := 2; i <= len(input); i += 3 {
		first, second, third := strings.Split(input[i-2], ""), strings.Split(input[i-1], ""), strings.Split(input[i], "")
		intersection := SliceIntersection(first, second, third)
		res += letterToValue(intersection)
	}
	return res
}

func letterToValue(input []string) int {
	var res int
	for _, val := range input {
		r := []rune(val)[0]
		if unicode.IsUpper(r) {
			res += int(r) - 38
		} else {
			res += int(r) - 96
		}
	}

	return res
}

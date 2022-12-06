package adventofcode2022

import (
	"strconv"
	"strings"
)

func dayFourChallengeOne(values []string) int {
	total := 0

	for _, pair := range values {

		parsedPair := strings.FieldsFunc(pair, func(r rune) bool {
			return r == '-' || r == ','
		})

		firstMin, _ := strconv.Atoi(parsedPair[0])
		firstMax, _ := strconv.Atoi(parsedPair[1])
		secondMin, _ := strconv.Atoi(parsedPair[2])
		secondMax, _ := strconv.Atoi(parsedPair[3])

		if (firstMin <= secondMin && firstMax >= secondMax) ||
			(secondMin <= firstMin && secondMax >= firstMax) {
			total += 1
		}
	}

	return total
}

func dayFourChallengeTwo(values []string) int {
	total := 0

	for _, pair := range values {

		parsedPair := strings.FieldsFunc(pair, func(r rune) bool {
			return r == '-' || r == ','
		})

		firstMin, _ := strconv.Atoi(parsedPair[0])
		firstMax, _ := strconv.Atoi(parsedPair[1])
		secondMin, _ := strconv.Atoi(parsedPair[2])
		secondMax, _ := strconv.Atoi(parsedPair[3])

		if (firstMax >= secondMin && firstMin <= secondMax) ||
			(secondMax >= firstMin && secondMin <= firstMax) {
			total += 1
		}
	}

	return total
}

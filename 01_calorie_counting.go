package adventofcode2022

import (
	"math"
	"sort"
	"strconv"
)

func dayOneChallengeOne(calories []string) int {
	max := math.MinInt
	counter := 0

	for _, strVal := range calories {

		if strVal == "" {
			max = Max(max, counter)
			counter = 0
			continue
		}

		val, _ := strconv.Atoi(strVal)
		counter += val
	}

	return max
}

func dayOneChallengeTwo(calories []string) int {
	top := []int{math.MinInt, math.MinInt, math.MinInt}
	counter := 0

	for _, strVal := range calories {

		if strVal == "" {
			top = append(top, counter)
			counter = 0
			continue
		}

		val, _ := strconv.Atoi(strVal)
		counter += val
	}

	sort.Ints(top)

	return top[len(top)-1] + top[len(top)-2] + top[len(top)-3]
}

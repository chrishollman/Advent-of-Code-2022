package adventofcode2022

func getFirstSubstringAllUnique(input string, count int) int {

	if len(input) < count {
		return -1
	}

	for i := count - 1; i < len(input); i++ {
		uniqueRunes := make(map[rune]bool, count)
		ss := input[i-count+1 : i+1]

		for _, r := range ss {
			uniqueRunes[r] = true
		}

		if len(uniqueRunes) == count {
			return i + 1
		}
	}

	return -1
}

func daySixChallengeOne(input string) int {
	return getFirstSubstringAllUnique(input, 4)
}

func daySixChallengeTwo(input string) int {
	return getFirstSubstringAllUnique(input, 14)
}

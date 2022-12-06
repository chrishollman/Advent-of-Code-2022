package adventofcode2022

func dayTwoChallengeOne(inputs []string) int {
	var total int
	for _, str := range inputs {
		total += calculateRoundExact(str)
	}

	return total
}

func dayTwoChallengeTwo(inputs []string) int {
	var total int
	for _, str := range inputs {
		total += calculateRoundStrategy(str)
	}

	return total
}

func calculateRoundExact(input string) int {
	var res int
	switch input {
	case "A X":
		res = 4 // Rock + Draw = 1 + 3
	case "A Y":
		res = 8 // Paper + Win = 2 + 6
	case "A Z":
		res = 3 // Scissors + Lose = 3 + 0
	case "B X":
		res = 1 // Rock + Lose = 1 + 0
	case "B Y":
		res = 5 // Paper + Draw = 2 + 3
	case "B Z":
		res = 9 // Scissors + Win = 3 + 6
	case "C X":
		res = 7 // Rock + Win = 1 + 6
	case "C Y":
		res = 2 // Paper + Loss = 2 + 0
	case "C Z":
		res = 6 // Scissors + Draw = 3 + 3
	}

	return res
}

func calculateRoundStrategy(input string) int {
	var res int
	switch input {
	case "A X":
		res = 3 // Scissors + Lose => 3 + 0
	case "A Y":
		res = 4 // Rock + Draw => 1 + 3
	case "A Z":
		res = 8 // Paper + Win => 2 + 6
	case "B X":
		res = 1 // Rock + Lose => 1 + 0
	case "B Y":
		res = 5 // Paper + Draw => 2 + 3
	case "B Z":
		res = 9 // Scissors + Win => 3 + 6
	case "C X":
		res = 2 // Paper + Lose => 2 + 0
	case "C Y":
		res = 6 // Scissors + Draw => 3 + 3
	case "C Z":
		res = 7 // Rock + Win => 1 + 6
	}

	return res
}

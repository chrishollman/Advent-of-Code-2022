package adventofcode2022

import (
	"fmt"
)

func dayEightChallengeOne(input []string) int {

	w := len(input[0])
	h := len(input)

	visible := make(map[string]bool, (w-1)*(h-1))
	visibleDefault := 2*w + 2*h - 4 // 4 edges - 4 corners

	// Top to Bottom
	for y := 1; y < h-1; y++ {

		// Left to Right
		largest := input[y][0]
		for x := 1; x < w-1; x++ {
			if largest == 57 {
				break
			}
			val := input[y][x]
			if val > largest {
				largest = val
				visible[fmt.Sprintf("%v,%v", y, x)] = true
				continue
			}
			if val == 57 {
				visible[fmt.Sprintf("%v,%v", y, x)] = true
				break
			}
		}

		// Right to Left
		largest = input[y][w-1]
		for x := w - 2; x > 0; x-- {
			if largest == 57 {
				break
			}
			val := input[y][x]
			if val > largest {
				largest = val
				visible[fmt.Sprintf("%v,%v", y, x)] = true
				continue
			}
			if val == 57 {
				visible[fmt.Sprintf("%v,%v", y, x)] = true
				break
			}
		}
	}

	// Left to Right
	for x := 1; x < w-1; x++ {

		// Top to Bottom
		largest := input[0][x]
		for y := 1; y < h-1; y++ {
			if largest == 57 {
				break
			}
			val := input[y][x]
			if val > largest {
				largest = val
				visible[fmt.Sprintf("%v,%v", y, x)] = true
				continue
			}
			if val == 57 {
				visible[fmt.Sprintf("%v,%v", y, x)] = true
				break
			}
		}

		// Bottom to Top
		largest = input[h-1][x]
		for y := h - 2; y > 0; y-- {
			if largest == 57 {
				break
			}
			val := input[y][x]
			if val > largest {
				largest = val
				visible[fmt.Sprintf("%v,%v", y, x)] = true
				continue
			}
			if val == 57 {
				visible[fmt.Sprintf("%v,%v", y, x)] = true
				break
			}
		}
	}

	return len(visible) + visibleDefault
}

func dayEightChallengeTwo(input []string) int {
	w := len(input[0])
	h := len(input)

	best := 0

	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			current := input[y][x]
			total := 1
			// Top

			canSee := 0
			for y2 := y - 1; y2 >= 0; y2-- {
				candidate := input[y2][x]
				canSee++
				if candidate >= current {
					break
				}

			}
			total *= canSee

			// Bottom
			canSee = 0
			for y2 := y + 1; y2 < h; y2++ {
				candidate := input[y2][x]
				canSee++
				if candidate >= current {
					break
				}
			}
			total *= canSee

			// Left
			canSee = 0
			for x2 := x - 1; x2 >= 0; x2-- {
				candidate := input[y][x2]
				canSee++
				if candidate >= current {
					break
				}
			}
			total *= canSee

			// Right
			canSee = 0
			for x2 := x + 1; x2 < w; x2++ {
				candidate := input[y][x2]
				canSee++
				if candidate >= current {
					break
				}
			}
			total *= canSee

			if total > best {
				best = total
			}
		}
	}

	return best
}

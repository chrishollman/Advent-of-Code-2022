package adventofcode2022

import (
	"bufio"
	"os"

	"golang.org/x/exp/constraints"
)

// Stack implementation lazily copied from https://www.educative.io/answers/how-to-implement-a-stack-in-golang

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true

}

type TestConfigStringSliceWantInt struct {
	input []string
	want  int
}

type TestConfigStringSliceWantString struct {
	input []string
	want  string
}

type TestConfigStringWantInt struct {
	input string
	want  int
}

func readFileIntoStringArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func MaxSlice[T constraints.Ordered](arg []T) T {
	if len(arg) == 0 {
		return *new(T)
	}

	max := arg[0]
	for _, val := range arg[1:] {
		if val > max {
			max = val
		}
	}

	return max
}

func MinSlice[T constraints.Ordered](arg []T) T {
	if len(arg) == 0 {
		return *new(T)
	}

	min := arg[0]
	for _, val := range arg[1:] {
		if val < min {
			min = val
		}
	}

	return min
}

func Max[T constraints.Ordered](args ...T) T {
	if len(args) == 0 {
		return *new(T)
	}

	max := args[0]
	for _, val := range args[1:] {
		if val > max {
			max = val
		}
	}

	return max
}

func Min[T constraints.Ordered](args ...T) T {
	if len(args) == 0 {
		return *new(T)
	}

	min := args[0]
	for _, val := range args[1:] {
		if val < min {
			min = val
		}
	}

	return min
}

// Generic slice intersection lazily copied from https://stackoverflow.com/a/72375462

func SliceIntersection[T constraints.Ordered](in ...[]T) []T {
	hash := make(map[T]*int)
	for _, slice := range in {
		duplicates := make(map[T]bool)
		for _, v := range slice {
			if _, ok := duplicates[v]; !ok {
				if counter := hash[v]; counter != nil {
					*counter++
				} else {
					i := 1
					hash[v] = &i
				}
				duplicates[v] = true
			}
		}
	}

	res := make([]T, 0)
	for v, c := range hash {
		if *c >= len(in) {
			res = append(res, v)
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

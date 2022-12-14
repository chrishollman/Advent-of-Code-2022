package adventofcode2022

import (
	"bufio"
	"os"

	"golang.org/x/exp/constraints"
)

// Stack implementation lazily copied from https://www.educative.io/answers/how-to-implement-a-stack-in-golang

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.vals) == 0 {
		return true
	}
	return false
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var blank T
		return blank, false
	}
	top := len(s.vals) - 1
	el := s.vals[top]
	s.vals = s.vals[:top]
	return el, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var blank T
		return blank, false
	}
	return s.vals[len(s.vals)-1], true
}

type Queue[T any] struct {
	vals []T
}

func (q *Queue[T]) IsEmpty() bool {
	if len(q.vals) == 0 {
		return true
	}
	return false
}

func (q *Queue[T]) Enqueue(val T) {
	q.vals = append(q.vals, val)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var blank T
		return blank, false
	}
	front := q.vals[0]
	q.vals = q.vals[1:]
	return front, true
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

// Example priority queue adapted from https://pkg.go.dev/container/heap

// Interfaces & associated ideas for A* from https://github.com/beefsack/go-astar
type pathNode interface {
	GetNeighbours() []pathNode
	EstCostToNeighbour(to pathNode) int
}

type pathNodeWrapper struct {
	pathNode pathNode
	cost     int
	rank     int
	parent   *pathNodeWrapper
	open     bool
	closed   bool
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*pathNodeWrapper

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].rank < pq[j].rank
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*pathNodeWrapper)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

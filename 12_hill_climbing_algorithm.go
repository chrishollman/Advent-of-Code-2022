package adventofcode2022

import (
	"container/heap"
	"fmt"
	"math"
)

type mazeNodeMap map[pathNode]*pathNodeWrapper

func (m mazeNodeMap) get(pNode pathNode) *pathNodeWrapper {
	target, ok := m[pNode]
	if !ok {
		target = &pathNodeWrapper{
			pathNode: pNode,
		}
		m[pNode] = target
	}
	return target
}

type maze map[int]map[int]*node

func (m maze) Node(x, y int) *node {
	if m[x] == nil {
		return nil
	}
	return m[x][y]
}

func (m maze) SetNode(mazeNode *node, elevation rune, value, x, y int) {
	if m[x] == nil {
		m[x] = map[int]*node{}
	}
	m[x][y] = mazeNode
	mazeNode.Elevation = elevation
	mazeNode.Value = value
	mazeNode.X = x
	mazeNode.Y = y
	mazeNode.Maze = m
}

func (m maze) FirstOfKind(v rune) *node {
	for _, row := range m {
		for _, t := range row {
			if t.Elevation == v {
				return t
			}
		}
	}
	return nil
}

func (m maze) AllOfKind(v rune) []*node {
	var res []*node
	for _, row := range m {
		for _, t := range row {
			if t.Elevation == v {
				res = append(res, t)
			}
		}
	}
	return res
}

func (m maze) Start() *node {
	return m.FirstOfKind('S')
}

func (m maze) End() *node {
	return m.FirstOfKind('E')
}

func (m maze) DrawSolution(solution []pathNode) {
	for y := 0; y < 41; y++ {
		for x := 0; x < 61; x++ {
			onPath := false
			for _, p := range solution {
				if p.(*node).X == x && p.(*node).Y == y {
					onPath = true
				}
			}
			if onPath {
				fmt.Printf("-")
			} else {
				fmt.Printf("%v", string(m[x][y].Elevation))
			}
		}
		fmt.Printf("\n")
	}
}

func buildMaze(input []string) maze {
	m := maze{}
	for y, row := range input {
		for x, r := range row {
			var v int
			switch r {
			case 'S':
				v = int('a') // 'S' == 'a' with respect to value
			case 'E':
				v = int('z') // 'E' == 'z' with respect to value
			default:
				v = int(r)
			}
			m.SetNode(&node{}, r, v, x, y)
		}
	}
	return m
}

type node struct {
	Elevation rune
	Value     int
	X, Y      int
	Maze      maze
}

func (n *node) GetNeighbours() []pathNode {
	neighbours := []pathNode{}
	for _, offset := range [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	} {
		if neighbour := n.Maze.Node(n.X+offset[0], n.Y+offset[1]); neighbour != nil && (neighbour.Value-n.Value <= 1) {
			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}

func (n *node) EstCostToNeighbour(to pathNode) int {
	toNode := to.(*node)
	absX := abs(toNode.X - n.X)
	absY := abs(toNode.Y - n.Y)
	return absX + absY
}

func dayTwelveChallengeOne(input []string) int {
	m := buildMaze(input)
	_, dist, success := Path(m.Start(), m.End())
	if !success {
		return -1
	}
	return dist
}

func dayTwelveChallengeTwo(input []string) int {
	min := math.MaxInt
	m := buildMaze(input)
	for _, from := range m.AllOfKind('a') {
		_, dist, success := Path(from, m.End())
		if success {
			min = Min(min, dist)
		}
	}
	return min
}

func Path(from, to *node) (path []pathNode, distance int, found bool) {
	nodeMap := mazeNodeMap{}
	prioQueue := &PriorityQueue{}
	heap.Init(prioQueue)

	fromNode := nodeMap.get(from)
	fromNode.open = true
	heap.Push(prioQueue, fromNode)

	for {
		if prioQueue.Len() == 0 {
			return
		}
		current := heap.Pop(prioQueue).(*pathNodeWrapper)
		current.open = false
		current.closed = true

		// Solution identified
		if current == nodeMap.get(to) {
			p := []pathNode{}
			curr := current
			for curr != nil {
				p = append(p, curr.pathNode)
				curr = curr.parent
			}
			return p, current.cost, true
		}

		// Gather paths
		for _, neighbour := range current.pathNode.GetNeighbours() {
			cost := current.cost + 1
			neighbourData := nodeMap.get(neighbour)
			if cost < neighbourData.cost {
				if neighbourData.open {
					heap.Remove(prioQueue, neighbourData.index)
				}
				neighbourData.open = false
				neighbourData.closed = false
			}
			if !neighbourData.open && !neighbourData.closed {
				neighbourData.cost = cost
				neighbourData.open = true
				neighbourData.rank = cost + neighbour.EstCostToNeighbour(to)
				neighbourData.parent = current
				heap.Push(prioQueue, neighbourData)
			}
		}
	}
}

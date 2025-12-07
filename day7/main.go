package main

import (
	"bytes"
	"os"
)

func part1(grid [][]byte, start [2]int) int {
	activeCols := make([]bool, len(grid[0]))
	activeCols[start[1]] = true
	totalSplits := 0
	for y := start[0] + 1; y < len(grid); y++ {
		row := grid[y]
		for x, cell := range row {
			if activeCols[x] && cell == '^' {
				activeCols[x-1] = true
				activeCols[x+1] = true
				activeCols[x] = false
				totalSplits++
			}
		}
	}
	return totalSplits
}

type node struct {
	left     *node
	right    *node
	numPaths uint64
}

func (n *node) calculateNumPaths() uint64 {
	if n.numPaths != 0 {
		return n.numPaths
	} else if n.left == nil && n.right == nil {
		n.numPaths = 1
		return 1
	}
	var total uint64
	total += n.left.calculateNumPaths()
	total += n.right.calculateNumPaths()
	n.numPaths = total
	return total
}

func (n *node) buildTree(grid [][]byte, start [2]int, cns map[[2]int]*node) {
	for y := start[0] + 1; y < len(grid); y++ {
		row := grid[y]
		cell := row[start[1]]
		if cell == '^' {
			if cn, ok := cns[[2]int{y, start[1] - 1}]; ok {
				n.left = cn
			} else {
				n.left = &node{}
				cns[[2]int{y, start[1] - 1}] = n.left
				n.left.buildTree(grid, [2]int{y, start[1] - 1}, cns)
			}
			if cn, ok := cns[[2]int{y, start[1] + 1}]; ok {
				n.right = cn
			} else {
				n.right = &node{}
				cns[[2]int{y, start[1] + 1}] = n.right
				n.right.buildTree(grid, [2]int{y, start[1] + 1}, cns)
			}
			return
		}
	}
}

func part2(grid [][]byte, start [2]int) uint64 {
	root := &node{}
	nodes := make(map[[2]int]*node)
	nodes[start] = root
	root.buildTree(grid, start, nodes)
	return root.calculateNumPaths()
}

func main() {
	data, _ := os.ReadFile("input.txt")
	grid := bytes.Split(data, []byte{'\n'})
	var start [2]int
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				start = [2]int{y, x}
			}
		}
	}
	println(part1(grid, start))
	println(part2(grid, start))
}

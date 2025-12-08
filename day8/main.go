package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

type junctionBox struct {
	x, y, z uint64
	circuit *circuit
}

type pair struct {
	a, b *junctionBox
	dist uint64
}

type circuit struct {
	size uint64
}

func manhattanDistance(a, b junctionBox) uint64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func main() {
	data, _ := os.ReadFile("input.txt")
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var coords []*junctionBox
	for scanner.Scan() {
		var c junctionBox
		_, _ = fmt.Sscanf(scanner.Text(), "%d,%d,%d", &c.x, &c.y, &c.z)
		coords = append(coords, &c)
	}
	var ps []pair
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			ps = append(ps, pair{a: coords[i], b: coords[j], dist: manhattanDistance(*coords[i], *coords[j])})
		}
	}
	slices.SortFunc(ps, func(a, b pair) int {
		return int(a.dist) - int(b.dist)
	})
	var orderedCircuits []*circuit
	var lastPair *pair
	for i := 0; i < len(ps); i++ {
		p := ps[i]
		if p.a.circuit != nil && p.a.circuit == p.b.circuit {
			continue
		}
		lastPair = &p
		if p.a.circuit == nil && p.b.circuit != nil {
			p.a.circuit = p.b.circuit
			p.a.circuit.size++
		} else if p.a.circuit != nil && p.b.circuit == nil {
			p.b.circuit = p.a.circuit
			p.b.circuit.size++
		} else if p.a.circuit == nil && p.b.circuit == nil {
			circuit := &circuit{size: 2}
			p.a.circuit = circuit
			p.b.circuit = circuit
			orderedCircuits = append(orderedCircuits, circuit)
		} else {
			circA := p.a.circuit
			circB := p.b.circuit
			totalSize := circA.size + circB.size
			for _, c := range coords {
				if c.circuit == circB {
					c.circuit = circA
				}
			}
			circA.size = totalSize
			for i, c := range orderedCircuits {
				if c == circB {
					orderedCircuits = append(orderedCircuits[:i], orderedCircuits[i+1:]...)
					break
				}
			}
		}

		if i == 999 {
			slices.SortFunc(orderedCircuits, func(a, b *circuit) int {
				return int(b.size) - int(a.size)
			})
			product := uint64(1)
			for i := 0; i < 3 && i < len(orderedCircuits); i++ {
				product *= orderedCircuits[i].size
			}
			println(product)
		}
	}

	println(lastPair.a.x * lastPair.b.x)
}

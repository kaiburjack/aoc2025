package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

type coordinate struct {
	x, y, z uint64
	circuit *circuit
}

type pair struct {
	a, b *coordinate
	dist uint64
}

type pairs []pair

func (p *pairs) Len() int           { return len(*p) }
func (p *pairs) Swap(i, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }
func (p *pairs) Less(i, j int) bool { return (*p)[i].dist < (*p)[j].dist }
func (p *pairs) Push(x interface{}) {
	*p = append(*p, x.(pair))
}
func (p *pairs) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

type circuit struct {
	size uint64
}

type circuits []*circuit

func (c *circuits) Len() int           { return len(*c) }
func (c *circuits) Swap(i, j int)      { (*c)[i], (*c)[j] = (*c)[j], (*c)[i] }
func (c *circuits) Less(i, j int) bool { return (*c)[i].size > (*c)[j].size }
func (c *circuits) Push(x interface{}) {
	*c = append(*c, x.(*circuit))
}
func (c *circuits) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

func manhattanDistance(a, b coordinate) uint64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func main() {
	data, _ := os.ReadFile("input.txt")
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var coords []*coordinate
	for scanner.Scan() {
		var c coordinate
		_, _ = fmt.Sscanf(scanner.Text(), "%d,%d,%d", &c.x, &c.y, &c.z)
		coords = append(coords, &c)
	}
	var ps pairs
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			ps = append(ps, pair{
				a:    coords[i],
				b:    coords[j],
				dist: manhattanDistance(*coords[i], *coords[j]),
			})
		}
	}
	heap.Init(&ps)
	var orderedCircuits circuits
	connections := 0
	var lastPair *pair
	for ps.Len() > 0 {
		p := heap.Pop(&ps).(pair)
		connections++
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

		if connections == 1000 {
			heap.Init(&orderedCircuits)
			var largestSizes []uint64
			for i := 0; i < 3 && orderedCircuits.Len() > 0; i++ {
				circuit := heap.Pop(&orderedCircuits).(*circuit)
				largestSizes = append(largestSizes, circuit.size)
			}
			product := uint64(1)
			for _, size := range largestSizes {
				product *= size
			}
			fmt.Printf("Product of largest circuit sizes: %d\n", product)
		}
	}

	xProduct := lastPair.a.x * lastPair.b.x
	fmt.Printf("Product of X coordinates of last connected pair: %d\n", xProduct)
}

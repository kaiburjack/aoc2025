package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type intRange struct {
	start uint64
	end   uint64
}
type rangeList []intRange

func (thiz *intRange) sizeOfOverlapWith(o intRange) uint64 {
	if thiz.end < o.start || o.end < thiz.start {
		return 0
	}
	return min(thiz.end, o.end) - min(thiz.start, o.start) + 1
}

func (thiz *intRange) isAdjacentTo(o intRange) bool {
	return thiz.end+1 == o.start || o.end+1 == thiz.start
}

func (thiz *rangeList) unionWith(o intRange) {
	var res = rangeList{o}
	for _, e := range *thiz {
		var merged bool
		for i := 0; i < len(res); i++ {
			r := &res[i]
			if r.sizeOfOverlapWith(e) > 0 || r.isAdjacentTo(e) {
				res[i].start = min(e.start, r.start)
				res[i].end = max(e.end, r.end)
				merged = true
				break
			}
		}
		if !merged {
			res = append(res, e)
		}
	}
	*thiz = res
}

func (thiz *rangeList) contains(value uint64) bool {
	for _, rg := range *thiz {
		if rg.start <= value && value <= rg.end {
			return true
		}
	}
	return false
}

func (thiz *rangeList) totalSize() uint64 {
	var total uint64
	for _, rg := range *thiz {
		total += rg.end - rg.start + 1
	}
	return total
}

func main() {
	data, _ := os.ReadFile("input.txt")
	array := bytes.Split(data, []byte{'\n'})
	var rs rangeList
	var i int
	var line []byte
	for i, line = range array {
		if len(line) == 0 {
			break
		}

		var a, b uint64
		_, _ = fmt.Sscanf(string(line), "%d-%d", &a, &b)
		rs.unionWith(intRange{start: a, end: b})
	}
	numFresh := 0
	for j := i + 1; j < len(array); j++ {
		p, _ := strconv.ParseUint(string(array[j]), 10, 64)
		if rs.contains(p) {
			numFresh++
		}
	}
	println("Part1:", numFresh)
	println("Part2:", rs.totalSize())
}

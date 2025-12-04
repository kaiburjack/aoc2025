package main

import (
	"bytes"
	"os"
)

func adjacent(array [][]byte, xc, yc int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			x, y := xc+dx, yc+dy
			if x < 0 || y < 0 || y >= len(array) || x >= len(array[y]) || x == xc && y == yc {
				continue
			}
			if array[y][x] == '@' {
				count++
			}
		}
	}
	return count
}

func part1(array [][]byte) int {
	count := 0
	for y := 0; y < len(array); y++ {
		for x := 0; x < len(array[y]); x++ {
			if array[y][x] == '@' && adjacent(array, x, y) < 4 {
				count++
			}
		}
	}
	return count
}

func part2(array [][]byte) int {
	count := 0
	for {
		removed := false
		for y := 0; y < len(array); y++ {
			for x := 0; x < len(array[y]); x++ {
				if array[y][x] == '@' && adjacent(array, x, y) < 4 {
					array[y][x] = '.'
					count++
					removed = true
				}
			}
		}
		if !removed {
			break
		}
	}
	return count
}

func main() {
	data, _ := os.ReadFile("input.txt")
	var array [][]byte
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		array = append(array, line)
	}

	println(part1(array))
	println(part2(array))
}

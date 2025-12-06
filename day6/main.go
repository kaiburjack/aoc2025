package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func calcNumbers(numbers []uint64, operator byte) uint64 {
	switch operator {
	case '+':
		var sum uint64
		for _, n := range numbers {
			sum += n
		}
		return sum
	case '*':
		product := uint64(1)
		for _, n := range numbers {
			product *= n
		}
		return product
	}
	return 0
}

func part1(grid [][]byte) uint64 {
	var numbers [][]uint64
	var operators []byte
	for _, line := range grid {
		if len(line) == 0 {
			break
		}
		parts := bytes.Fields(line)
		for i, part := range parts {
			num, err := strconv.ParseUint(string(part), 10, 64)
			if err != nil {
				if len(operators) <= i {
					operators = append(operators, part[0])
				}
			} else {
				if len(numbers) <= i {
					numbers = append(numbers, []uint64{})
				}
				numbers[i] = append(numbers[i], num)
			}
		}
	}
	var total uint64
	for i, nums := range numbers {
		total += calcNumbers(nums, operators[i])
	}
	return total
}

func part2(grid [][]byte) uint64 {
	var numbers []uint64
	var total uint64
	for col := len(grid[0]) - 1; col >= 0; col-- {
		var number uint64
		for row := 0; row < len(grid); row++ {
			c := grid[row][col]
			if c >= '0' && c <= '9' {
				number = number*10 + uint64(c-'0')
			}
			if c == '+' || c == '*' {
				numbers = append(numbers, number)
				total += calcNumbers(numbers, c)
				numbers = numbers[:0]
			} else if row == len(grid)-1 {
				if number != 0 {
					numbers = append(numbers, number)
				}
			}
		}
	}
	return total
}

func padGrid(grid [][]byte) {
	maxLen := 0
	for _, row := range grid {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}
	for i := range grid {
		for len(grid[i]) < maxLen {
			grid[i] = append(grid[i], ' ')
		}
	}
}

func main() {
	data, _ := os.ReadFile("input.txt")
	grid := bytes.Split(data, []byte{'\n'})
	padGrid(grid)
	fmt.Println("Part 1:", part1(grid))
	fmt.Println("Part 2:", part2(grid))
}

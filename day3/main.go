package main

import (
	"bytes"
	"os"
	"time"
)

func find(s string, acc, remaining int) int {
	if remaining == 0 {
		return acc
	}
	best, j := -1, 0
	for i := 0; i < len(s)-remaining+1; i++ {
		digit := int(s[i] - '0')
		if digit > best {
			best = digit
			j = i
		}
	}
	return find(s[j+1:], acc*10+best, remaining-1)
}

func main() {
	data, _ := os.ReadFile("input.txt")
	part1, part2 := 0, 0

	start := time.Now()
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		part1 += find(string(line), 0, 2)
		part2 += find(string(line), 0, 12)
	}
	end := time.Now()

	println("Execution Time:", end.Sub(start).String())
	println(part1)
	println(part2)
}

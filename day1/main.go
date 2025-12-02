package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	scanner := bufio.NewScanner(bytes.NewReader(input))
	pos, part1, part2 := 50, 0, 0
	time1 := time.Now()
	for scanner.Scan() {
		buf := scanner.Bytes()
		d, steps := buf[0], buf[1:]
		n, _ := strconv.Atoi(string(steps))
		if d == 'L' {
			for i := 0; i < n; i++ {
				pos = (pos - 1) % 100
				if pos == 0 {
					part2++
				}
			}
		} else if d == 'R' {
			for i := 0; i < n; i++ {
				pos = (pos + 1) % 100
				if pos == 0 {
					part2++
				}
			}
		}
		if pos == 0 {
			part1++
		}
	}
	time2 := time.Now()
	fmt.Println("Execution Time:", time2.Sub(time1).String())
	fmt.Println(part1)
	fmt.Println(part2)
}

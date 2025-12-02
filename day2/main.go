package main

import (
	"bytes"
	"os"
	"strconv"
	"time"
)

func containsOnlyRepeatedSubstring(s string) (bool, int) {
	for i := len(s) / 2; i >= 1; i-- {
		if len(s)%i != 0 {
			continue
		}
		invalid := true
		for j := i; j < len(s); j += i {
			if s[j:j+i] != s[0:i] {
				invalid = false
				break
			}
		}
		if invalid {
			return true, i
		}
	}
	return false, 0
}

const fileName = "input.txt"

func main() {
	data, _ := os.ReadFile(fileName)
	start := time.Now()
	part1, part2 := 0, 0
	for _, part := range bytes.Split(data, []byte(",")) {
		r := bytes.Split(part, []byte("-"))
		e, _ := strconv.Atoi(string(r[1]))
		for i, _ := strconv.Atoi(string(r[0])); i <= e; i++ {
			n := strconv.Itoa(i)
			if invalid, k := containsOnlyRepeatedSubstring(n); invalid {
				if k == len(n)/2 {
					part1 += i
				}
				part2 += i
			}
		}
	}
	end := time.Now()
	println("Execution Time:", end.Sub(start).String())
	println(part1)
	println(part2)
}

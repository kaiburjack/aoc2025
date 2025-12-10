package main

import (
	"bufio"
	"math"
	"os"

	"github.com/alecthomas/participle/v2"
)

type ParsedButton struct {
	Indexes []int `"(" @Int (',' @Int)* ")"`
}

type ParsedLine struct {
	Mask     string         `"[" @('#' | '.')+ "]"`
	Buttons  []ParsedButton `@@+`
	Joltages []int          `"{" @Int (',' @Int)* "}"`
}

type Line struct {
	Mask     int
	Buttons  []int
	Joltages []int
}

func (pl *ParsedLine) toLine() Line {
	var line Line
	for i, c := range pl.Mask {
		if c == '#' {
			line.Mask |= 1 << i
		}
	}
	line.Buttons = make([]int, 0, len(pl.Buttons))
	for _, pb := range pl.Buttons {
		var button int
		for _, index := range pb.Indexes {
			button |= 1 << index
		}
		line.Buttons = append(line.Buttons, button)
	}
	line.Joltages = pl.Joltages
	return line
}

func (l *Line) part1Rec(p, b, m int) int {
	if m == l.Mask {
		return p
	} else if b >= len(l.Buttons) {
		return math.MaxInt
	}
	wo := l.part1Rec(p, b+1, m)
	w := l.part1Rec(p+1, b+1, m^l.Buttons[b])
	if w < wo {
		wo = w
	}
	return wo
}

func (l *Line) part1() int {
	return l.part1Rec(0, 0, 0)
}

func main() {
	parser, _ := participle.Build[ParsedLine]()
	file, _ := os.Open("input.txt")
	s := bufio.NewScanner(file)
	part1 := 0
	for s.Scan() {
		parsedLine, _ := parser.ParseString("", s.Text())
		line := (*parsedLine).toLine()
		part1 += line.part1()
	}
	println("Part 1:", part1)
}

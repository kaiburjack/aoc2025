package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type pt struct {
	x, y int64
}

type line struct {
	a, b pt
}

type rect struct {
	min, max pt
	area     int64
}

func A(min, max pt) int64 {
	return (max.x - min.x + 1) * (max.y - min.y + 1)
}

func li(l1, l2 line) bool {
	if l1.a.x == l1.b.x && l2.a.y == l2.b.y {
		if l1.a.x > min(l2.a.x, l2.b.x) &&
			l1.a.x < max(l2.a.x, l2.b.x) &&
			l2.a.y > min(l1.a.y, l1.b.y) &&
			l2.a.y < max(l1.a.y, l1.b.y) {
			return true
		}
	} else if l1.a.y == l1.b.y && l2.a.x == l2.b.x {
		if l2.a.x > min(l1.a.x, l1.b.x) &&
			l2.a.x < max(l1.a.x, l1.b.x) &&
			l1.a.y > min(l2.a.y, l2.b.y) &&
			l1.a.y < max(l2.a.y, l2.b.y) {
			return true
		}
	}
	return false
}

func rip(r rect, lines []line) bool {
	cs := []pt{
		{r.min.x, r.min.y},
		{r.min.x, r.max.y},
		{r.max.x, r.min.y},
		{r.max.x, r.max.y},
	}
	for _, c := range cs {
		if !pip(c, lines) {
			return false
		}
	}
	edges := []line{
		{a: pt{r.min.x, r.min.y}, b: pt{r.max.x, r.min.y}},
		{a: pt{r.max.x, r.min.y}, b: pt{r.max.x, r.max.y}},
		{a: pt{r.max.x, r.max.y}, b: pt{r.min.x, r.max.y}},
		{a: pt{r.min.x, r.max.y}, b: pt{r.min.x, r.min.y}},
	}
	for _, e1 := range edges {
		for _, e2 := range lines {
			if li(e1, e2) {
				return false
			}
		}
	}
	return true
}

func pip(p pt, lines []line) bool {
	in := false
	for _, l := range lines {
		if l.a.y == l.b.y {
			if p.y == l.a.y &&
				p.x >= min(l.a.x, l.b.x) &&
				p.x <= max(l.a.x, l.b.x) {
				return true
			}
			if p.y < l.a.y &&
				(l.a.x > p.x && l.b.x <= p.x ||
					l.b.x > p.x && l.a.x <= p.x) {
				in = !in
			}
		} else if p.x == l.a.x &&
			p.y >= min(l.a.y, l.b.y) &&
			p.y <= max(l.a.y, l.b.y) {
			return true
		}
	}
	return in
}

func lines(pts []pt) []line {
	var ls []line
	for i := 0; i < len(pts); i++ {
		next := (i + 1) % len(pts)
		ls = append(ls, line{a: pts[i], b: pts[next]})
	}
	return ls
}

func part1(pts []pt) (int64, []rect) {
	var rects []rect
	var maxA int64
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			r := rect{
				min: pt{
					x: min(pts[i].x, pts[j].x),
					y: min(pts[i].y, pts[j].y),
				},
				max: pt{
					x: max(pts[i].x, pts[j].x),
					y: max(pts[i].y, pts[j].y),
				},
			}
			r.area = A(r.min, r.max)
			rects = append(rects, r)
			if r.area > maxA {
				maxA = r.area
			}
		}
	}
	return maxA, rects
}

func part2(pts []pt, rects []rect) int64 {
	ls := lines(pts)
	var maxA int64
	for _, r := range rects {
		if r.area > maxA && rip(r, ls) {
			maxA = r.area
		}
	}
	return maxA
}

func main() {
	data, _ := os.ReadFile("input.txt")
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var pts []pt
	for scanner.Scan() {
		var p pt
		_, _ = fmt.Sscanf(scanner.Text(), "%d,%d", &p.x, &p.y)
		pts = append(pts, p)
	}
	A, rects := part1(pts)
	fmt.Println("Part1:", A)
	fmt.Println("Part2:", part2(pts, rects))
}

package y2016

import (
	"math"
	"strings"
)

type Solver06 struct{}
type columnChars []map[rune]int

func countColumnChars(lines []string) columnChars {
	arrs := make([]map[rune]int, len(lines[0]))
	for i := range arrs {
		arrs[i] = make(map[rune]int)
	}

	for _, line := range lines {
		for i, r := range line {
			arrs[i][r]++
		}
	}
	return arrs
}

func (s *Solver06) SolveP1(in string) string {
	lines := strings.Split(in, "\n")
	chars := countColumnChars(lines)

	var msg string
	for _, v := range chars {
		var max int
		var r rune
		for k, v := range v {
			if max < v {
				r = k
				max = v
			}
		}
		msg += string(r)
	}

	return msg
}

func (s *Solver06) SolveP2(in string) string {
	lines := strings.Split(in, "\n")
	chars := countColumnChars(lines)

	var msg string
	for _, v := range chars {
		min := math.MaxInt
		var r rune
		for k, v := range v {
			if v < min {
				r = k
				min = v
			}
		}
		msg += string(r)
	}

	return msg
}

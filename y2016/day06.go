package y2016

import (
	"bufio"
	"context"
	"io"
	"math"
)

type Solver06 struct{}
type columnChars []map[rune]int

func countColumnChars(ch <-chan string) columnChars {
	init := true
	var arrs columnChars

	for v := range ch {
		if init {
			arrs = make([]map[rune]int, len(v))
			for i := range arrs {
				arrs[i] = make(map[rune]int)
			}
			init = false
		}

		for i, r := range v {
			arrs[i][r]++
		}
	}

	return arrs
}

func (s *Solver06) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	run := func(chars columnChars) string {
		var str string
		for _, v := range chars {
			var max int
			var r rune
			for k, v := range v {
				if max < v {
					r = k
					max = v
				}
			}
			str += string(r)
		}
		return str
	}

	return s.solve(ctx, r, run)
}

func (s *Solver06) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	run := func(chars columnChars) string {
		var str string
		for _, v := range chars {
			min := math.MaxInt
			var r rune
			for k, v := range v {
				if v < min {
					r = k
					min = v
				}
			}
			str += string(r)
		}
		return str
	}

	return s.solve(ctx, r, run)
}

func (s *Solver06) solve(ctx context.Context, r io.Reader, run func(columnChars) string) (string, error) {
	ch := make(chan string)
	go scan(r, ch, bufio.ScanLines)
	chars := countColumnChars(ch)
	return run(chars), nil
}

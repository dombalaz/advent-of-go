package y2016

import (
	"bufio"
	"context"
	"io"
	"strconv"
	"strings"
)

func isValidTriangle(a, b, c int) bool {
	if a < b+c && b < a+c && c < a+b {
		return true
	}
	return false
}

type Solver03 struct{}

func countTriangles(ch <-chan string) int {
	valid := 0
	for v := range ch {
		sides := strings.Fields(v)

		a, _ := strconv.Atoi(sides[0])
		b, _ := strconv.Atoi(sides[1])
		c, _ := strconv.Atoi(sides[2])

		if isValidTriangle(a, b, c) {
			valid += 1
		}
	}
	return valid
}

func (s *Solver03) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	ch := make(chan string)
	var err error
	go func() {
		err = scan(r, ch, bufio.ScanLines)
	}()
	return strconv.FormatInt(int64(countTriangles(ch)), 10), err
}

func (s *Solver03) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// TODO(Dominik Balaz): Try to write scan function for this
	rotateLines := func() {
		defer close(ch2)
		var rotatedLines [3]string
		var i int

		for v := range ch1 {
			sides := strings.Fields(v)
			for i, v := range sides {
				rotatedLines[i] += v + " "
			}

			i++
			if i == 3 {
				for _, v := range rotatedLines {
					ch2 <- v
				}
				i = 0
				rotatedLines = [3]string{}
			}
		}
	}

	var err error
	go func() {
		err = scan(r, ch1, bufio.ScanLines)
	}()
	go rotateLines()
	return strconv.FormatInt(int64(countTriangles(ch2)), 10), err
}

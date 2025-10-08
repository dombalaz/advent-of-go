package y2016_test

import (
	"strings"
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestLcdScreenOperation(t *testing.T) {
	ops := []string{
		"rect 3x2",
		"rotate column x=1 by 1",
		"rotate row y=0 by 4",
		"rotate column x=1 by 1",
	}
	w := "6"
	solver := y2016.CreateSolver(y2016.Day08)

	for i := range len(ops) {
		in := strings.Join(ops[:i+1], "\n")
		r := solver.SolveP1(in)
		if r != w {
			t.Errorf("solver.SolveP1(%v) = %v, want %v", in, r, w)
		}
	}
}

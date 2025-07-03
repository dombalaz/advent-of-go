package y2016_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestKeypadInstructions(t *testing.T) {
	in :=
		`ULL
RRDDD
LURDL
UUUUD`

	solver := y2016.CreateSolver(y2016.Day02)
	r := solver.SolveP1(in)
	w := "1985"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = %v, want %v", in, r, w)
	}

	in = "RRR"
	r = solver.SolveP1(in)
	w = "6"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = %v, want %v", in, r, w)
	}
}

func TestHexKeypadInstructions(t *testing.T) {
	in :=
		`ULL
RRDDD
LURDL
UUUUD`

	solver := y2016.CreateSolver(y2016.Day02)
	r := solver.SolveP2(in)
	w := "5DB3"
	if r != w {
		t.Errorf("solver.SolveP2(%v) = %v, want %v", in, r, w)
	}
}

package y2016_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestValidTriangle(t *testing.T) {
	solver := y2016.CreateSolver(y2016.Day03)
	in := "5 6 7"
	v := solver.SolveP1(in)
	w := "1"
	if v != w {
		t.Errorf("solver.SolveP1(%v) = %v, want %v", in, v, w)
	}
}

func TestInvalidTriangle(t *testing.T) {
	solver := y2016.CreateSolver(y2016.Day03)
	in := "5 10 25"
	v := solver.SolveP1(in)
	w := "0"
	if v != w {
		t.Errorf("solver.SolveP1(%v) = %v, want %v", in, v, w)
	}
}

func TestMultipleSpaces(t *testing.T) {
	solver := y2016.CreateSolver(y2016.Day03)
	in := "5   6   7"
	v := solver.SolveP1(in)
	w := "1"
	if v != w {
		t.Errorf("solver.SolveP1(%v) = %v, want %v", in, v, w)
	}
}

func TestValidVerticalTriangle(t *testing.T) {
	in := `101 301 501
102 302 502
103 303 503`

	solver := y2016.CreateSolver(y2016.Day03)
	v := solver.SolveP2(in)
	w := "3"
	if v != w {
		t.Errorf("solver.SolveP2(%v) = %v, want %v", in, v, w)
	}
}

func TestInvalidVerticalTriangle(t *testing.T) {
	in := `201 401 601
202 402 602
203 403 603`

	solver := y2016.CreateSolver(y2016.Day03)
	v := solver.SolveP2(in)
	w := "3"
	if v != w {
		t.Errorf("solver.SolveP2(%v) = %v, want %v", in, v, w)
	}
}

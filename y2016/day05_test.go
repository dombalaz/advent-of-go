package y2016_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestMD5PasswordGen(t *testing.T) {
	id := "abc"
	solver := y2016.CreateSolver(y2016.Day05)
	r := solver.SolveP1(id)
	w := "18f47a30"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = %v, want %v", id, r, w)
	}
}

func TestMD5PasswordGen2(t *testing.T) {
	id := "abc"
	solver := y2016.CreateSolver(y2016.Day05)
	r := solver.SolveP2(id)
	w := "05ace8e3"
	if r != w {
		t.Errorf("solver.SolveP2(%v) = %v, want %v", id, r, w)
	}
}

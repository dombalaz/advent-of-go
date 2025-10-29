package y2016_test

import (
	"context"
	"strings"
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestMD5PasswordGen(t *testing.T) {
	id := "abc"
	solver := y2016.Solver05{}
	r, _ := solver.SolveP1(context.Background(), strings.NewReader(id))
	w := "18f47a30"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = %v, want %v", id, r, w)
	}
}

func TestMD5PasswordGen2(t *testing.T) {
	id := "abc"
	solver := y2016.Solver05{}
	r, _ := solver.SolveP2(context.Background(), strings.NewReader(id))
	w := "05ace8e3"
	if r != w {
		t.Errorf("solver.SolveP2(%v) = %v, want %v", id, r, w)
	}
}

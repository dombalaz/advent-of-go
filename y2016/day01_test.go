package y2016_test

import (
	"context"
	"strings"
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestAbsNegative(t *testing.T) {
	a := -1
	r := y2016.Abs(a)
	w := a * (-1)
	if r != w {
		t.Errorf("abs(%v) = %v, want %v", a, r, w)
	}
}

func TestAbsZero(t *testing.T) {
	a := 0
	r := y2016.Abs(a)
	w := a * (-1)
	if r != w {
		t.Errorf("abs(%v) = %v, want %v", a, r, w)
	}
}

func TestAbsPositive(t *testing.T) {
	a := 1
	r := y2016.Abs(a)
	if r != a {
		t.Errorf("abs(%v) = %v, want %v", a, r, a)
	}
}

func TestGridDistance(t *testing.T) {
	cases := []struct {
		in string
		w  string
	}{{in: "R2, L3", w: "5"}, {in: "R2, R2, R2", w: "2"}, {in: "R5, L5, R5, R3", w: "12"}}

	solver := y2016.Solver01{}
	for _, v := range cases {
		d, _ := solver.SolveP1(context.Background(), strings.NewReader(v.in))
		if d != v.w {
			t.Errorf("solver.SolveP1(%v) = %v, want %v", v.in, d, v.w)
		}
	}
}

func TestTwiceVisitDistance(t *testing.T) {
	in := "R8, R4, R4, R8"
	solver := y2016.Solver01{}
	d, _ := solver.SolveP2(context.Background(), strings.NewReader(in))
	w := "4"
	if d != w {
		t.Errorf("solver.SolveP2(%v) = %v, want %v", in, d, w)
	}
}

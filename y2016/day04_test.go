package y2016_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestRealRooms(t *testing.T) {
	in := "aaaaa-bbb-z-y-x-123[abxyz]"
	solver := y2016.CreateSolver(y2016.Day04)
	r := solver.SolveP1(in)
	w := "123"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = '%v', want '%v'", in, r, w)
	}

	in = "a-b-c-d-e-f-g-h-987[abcde]"
	r = solver.SolveP1(in)
	w = "987"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = '%v', want '%v'", in, r, w)
	}

	in = "not-a-real-room-404[oarel]"
	r = solver.SolveP1(in)
	w = "404"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = '%v', want '%v'", in, r, w)
	}

	in = "totally-real-room-200[decoy]"
	r = solver.SolveP1(in)
	w = "0"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = '%v', want '%v'", in, r, w)
	}
}

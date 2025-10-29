package y2016_test

import (
	"context"
	"strings"
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestRealRooms(t *testing.T) {
	in := "aaaaa-bbb-z-y-x-123[abxyz]"
	solver := y2016.Solver04{}
	r, _ := solver.SolveP1(context.Background(), strings.NewReader(in))
	w := "123"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = '%v', want '%v'", in, r, w)
	}

	in = "a-b-c-d-e-f-g-h-987[abcde]"
	r, _ = solver.SolveP1(context.Background(), strings.NewReader(in))
	w = "987"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = '%v', want '%v'", in, r, w)
	}

	in = "not-a-real-room-404[oarel]"
	r, _ = solver.SolveP1(context.Background(), strings.NewReader(in))
	w = "404"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = '%v', want '%v'", in, r, w)
	}

	in = "totally-real-room-200[decoy]"
	r, _ = solver.SolveP1(context.Background(), strings.NewReader(in))
	w = "0"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = '%v', want '%v'", in, r, w)
	}
}

func TestNorthPoleObjectStorage(t *testing.T) {
	in := "ghkma-7[aghkm]"
	solver := y2016.Solver04{}
	r, _ := solver.SolveP2(context.Background(), strings.NewReader(in))
	w := "7"
	if r != w {
		t.Errorf("solver.SolveP2(%v) = '%v', want '%v'", in, r, w)
	}
}

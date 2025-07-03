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

	r, _ := y2016.Solve2(in)
	w := "1985"
	if r != w {
		t.Errorf("y2016.Solve2(%v) = %v, want %v", in, r, w)
	}

	in = "RRR"
	r, _ = y2016.Solve2(in)
	w = "6"
	if r != w {
		t.Errorf("y2016.Solve2(%v) = %v, want %v", in, r, w)
	}
}

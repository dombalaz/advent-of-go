package y2016_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestErrorCorrectedMessage(t *testing.T) {
	in := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

	solver := y2016.CreateSolver(y2016.Day06)
	r := solver.SolveP1(in)
	w := "easter"
	if r != w {
		t.Errorf("solver.SolveP1(%v) = %v, want %v", in, r, w)
	}
}

func TestErrorCorrectedMessage2(t *testing.T) {
	in := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

	solver := y2016.CreateSolver(y2016.Day06)
	r := solver.SolveP2(in)
	w := "advent"
	if r != w {
		t.Errorf("solver.SolveP2(%v) = %v, want %v", in, r, w)
	}
}

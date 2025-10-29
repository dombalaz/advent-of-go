package y2016_test

import (
	"context"
	"strings"
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

	solver := y2016.Solver06{}
	r, _ := solver.SolveP1(context.Background(), strings.NewReader(in))
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

	solver := y2016.Solver06{}
	r, _ := solver.SolveP2(context.Background(), strings.NewReader(in))
	w := "advent"
	if r != w {
		t.Errorf("solver.SolveP2(%v) = %v, want %v", in, r, w)
	}
}

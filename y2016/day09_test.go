package y2016_test

import (
	"strings"
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestDetermineSize(t *testing.T) {
	tests := []struct {
		in string
		w  int
		v2 bool
	}{
		{
			in: "ADVENT",
			w:  6,
		},
		{
			in: "A(1x5)BC",
			w:  7,
		},
		{
			in: "(3x3)XYZ",
			w:  9,
		},
		{
			in: "A(2x2)BCD(2x2)EFG",
			w:  11,
		},
		{
			in: "(6x1)(1x3)A",
			w:  6,
		},
		{
			in: "X(8x2)(3x3)ABCY",
			w:  18,
		},
		{
			in: "(3x3)XYZ",
			w:  9,
			v2: true,
		},
		{
			in: "X(8x2)(3x3)ABCY",
			w:  20,
			v2: true,
		},
		{
			in: "(27x12)(20x12)(13x14)(7x10)(1x12)A",
			w:  241920,
			v2: true,
		},
		{
			in: "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
			w:  445,
			v2: true,
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			r, err := y2016.DetermineSize(strings.NewReader(test.in), test.v2)
			if err != nil {
				t.Errorf("y2016.DetermineSize(%v, %v) = %v, want nil", test.in, test.v2, err)
			}
			if r != test.w {
				t.Errorf("y2016.DetermineSize(%v, %v) = %v, want %v", test.in, test.v2, r, test.w)
			}
		})
	}
}

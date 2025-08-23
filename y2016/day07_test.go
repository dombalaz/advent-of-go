package y2016_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestIPv7TLSSupport(t *testing.T) {
	tests := []struct {
		in string
		w  string
	}{
		{
			in: "abba[mnop]qrst",
			w:  "1",
		},
		{
			in: "abcd[bddb]xyyx",
			w:  "0",
		},
		{
			in: "aaaa[qwer]tyui",
			w:  "0",
		},
		{
			in: "ioxxoj[asdfgh]zxcvbn",
			w:  "1",
		},
	}

	solver := y2016.CreateSolver(y2016.Day07)
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			r := solver.SolveP1(test.in)
			if r != test.w {
				t.Errorf("solver.SolveP1(%v) = %v, want %v", test.in, r, test.w)
			}
		})
	}
}

func TestIPv7SSLSupport(t *testing.T) {
	tests := []struct {
		in string
		w  string
	}{
		{
			in: "aba[bab]xyz",
			w:  "1",
		},
		{
			in: "xyx[xyx]xyx",
			w:  "0",
		},
		{
			in: "aaa[kek]eke",
			w:  "1",
		},
		{
			in: "zazbz[bzb]cdb",
			w:  "1",
		},
	}

	solver := y2016.CreateSolver(y2016.Day07)
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			r := solver.SolveP2(test.in)
			if r != test.w {
				t.Errorf("solver.SolveP1(%v) = %v, want %v", test.in, r, test.w)
			}
		})
	}
}

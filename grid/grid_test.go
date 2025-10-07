package grid_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/grid"
)

func TestNew(t *testing.T) {
	x := 3
	y := 3
	g := grid.New[int](x, y)
	for i := range x {
		for j := range y {
			at := g.At(i, j)
			if at != 0 {
				t.Errorf("g.At(%v, %v) = %v, want %v", i, j, at, 0)
			}

		}
	}
}

func TestDimensions(t *testing.T) {
	x := 2
	y := 3
	g := grid.New[int](x, y)
	if g.X() != x {
		t.Errorf("g.X() = %v, want %v", g.X(), x)
	}

	if g.Y() != y {
		t.Errorf("g.Y() = %v, want %v", g.Y(), x)
	}
}

func TestFill(t *testing.T) {
	x := 2
	y := 3
	g := grid.New[int](x, y)
	slice := []int{1, 2, 3, 4, 5}
	g.Fill(slice)
	// Expects grid
	// 5 0
	// 3 4
	// 1 2
	tests := []struct {
		x int
		y int
		w int
	}{
		{x: 0, y: 0, w: 1},
		{x: 1, y: 0, w: 2},
		{x: 0, y: 1, w: 3},
		{x: 1, y: 1, w: 4},
		{x: 0, y: 2, w: 5},
		{x: 1, y: 2, w: 0},
	}
	for _, test := range tests {
		at := g.At(test.x, test.y)
		if at != test.w {
			t.Errorf("g.At(%v, %v) = %v, want %v", test.x, test.y, at, test.w)
		}
	}
}

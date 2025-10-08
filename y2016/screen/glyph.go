package screen

import (
	"errors"

	"github.com/dombalaz/advent-of-go/grid"
)

const glyphHeight = 6
const glyphWidth = 5

type Glyph grid.Grid[bool]

func NewGlyph() Glyph {
	return Glyph(grid.New[bool](glyphWidth, glyphHeight))
}

func NewGlyphFromRune(r rune) (Glyph, error) {
	switch r {
	case 'B':
		return glyphB, nil
	case 'C':
		return glyphC, nil
	case 'E':
		return glyphE, nil
	case 'F':
		return glyphF, nil
	case 'J':
		return glyphJ, nil
	case 'L':
		return glyphL, nil
	case 'O':
		return glyphO, nil
	case 'P':
		return glyphP, nil
	case 'U':
		return glyphU, nil
	case 'Z':
		return glyphZ, nil
	}
	return NewGlyph(), errors.New("cannot convert rune to glyph")
}

// ###..
// #..#.
// ###..
// #..#.
// #..#.
// ###..
var glyphB = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	true, true, true, false, false,
	true, false, false, true, false,
	true, true, true, false, false,
	true, false, false, true, false,
	true, false, false, true, false,
	true, true, true, false, false,
}))

// .##..
// #..#.
// #....
// #....
// #..#.
// .##..
var glyphC = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	false, true, true, false, false,
	true, false, false, true, false,
	true, false, false, false, false,
	true, false, false, false, false,
	true, false, false, true, false,
	false, true, true, false, false,
}))

// ####.
// #....
// ###..
// #....
// #....
// ####.
var glyphE = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	true, true, true, true, false,
	true, false, false, false, false,
	true, true, true, false, false,
	true, false, false, false, false,
	true, false, false, false, false,
	true, true, true, true, false,
}))

// ####.
// #....
// ###..
// #....
// #....
// #....
var glyphF = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	true, true, true, true, false,
	true, false, false, false, false,
	true, true, true, false, false,
	true, false, false, false, false,
	true, false, false, false, false,
	true, false, false, false, false,
}))

// ..##.
// ...#.
// ...#.
// ...#.
// #..#.
// .##..
var glyphJ = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	false, false, true, true, false,
	false, false, false, true, false,
	false, false, false, true, false,
	false, false, false, true, false,
	true, false, false, true, false,
	false, true, true, false, false,
}))

// #....
// #....
// #....
// #....
// #....
// ####.
var glyphL = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	true, false, false, false, false,
	true, false, false, false, false,
	true, false, false, false, false,
	true, false, false, false, false,
	true, false, false, false, false,
	true, true, true, true, false,
}))

// .##..
// #..#.
// #..#.
// #..#.
// #..#.
// .##..
var glyphO = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	false, true, true, false, false,
	true, false, false, true, false,
	true, false, false, true, false,
	true, false, false, true, false,
	true, false, false, true, false,
	false, true, true, false, false,
}))

// ###..
// #..#.
// #..#.
// ###..
// #....
// #....
var glyphP = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	true, true, true, false, false,
	true, false, false, true, false,
	true, false, false, true, false,
	true, true, true, false, false,
	true, false, false, false, false,
	true, false, false, false, false,
}))

// #..#.
// #..#.
// #..#.
// #..#.
// #..#.
// .##..
var glyphU = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	true, false, false, true, false,
	true, false, false, true, false,
	true, false, false, true, false,
	true, false, false, true, false,
	true, false, false, true, false,
	false, true, true, false, false,
}))

// ####.
// ...#.
// ..#..
// .#...
// #....
// ####.
var glyphZ = Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, []bool{
	true, true, true, true, false,
	false, false, false, true, false,
	false, false, true, false, false,
	false, true, false, false, false,
	true, false, false, false, false,
	true, true, true, true, false,
}))

func (g *Glyph) String() string {
	cnv := (*grid.Grid[bool])(g)
	var r string
	for y := range cnv.Y() {
		if y != 0 {
			r += "\n"
		}
		for x := range cnv.X() {
			if cnv.At(x, y) {
				r += "#"
			} else {
				r += "."
			}
		}
	}
	return r
}

func (g *Glyph) Rune() (rune, error) {
	switch g.String() {
	case glyphB.String():
		return 'B', nil
	case glyphC.String():
		return 'C', nil
	case glyphE.String():
		return 'E', nil
	case glyphF.String():
		return 'F', nil
	case glyphJ.String():
		return 'J', nil
	case glyphL.String():
		return 'L', nil
	case glyphO.String():
		return 'O', nil
	case glyphP.String():
		return 'P', nil
	case glyphU.String():
		return 'U', nil
	case glyphZ.String():
		return 'Z', nil
	}
	return 0, errors.New("cannot convert to rune")
}

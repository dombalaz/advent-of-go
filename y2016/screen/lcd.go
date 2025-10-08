package screen

import (
	"strconv"
	"strings"

	"github.com/dombalaz/advent-of-go/grid"
)

type LCD struct {
	g grid.Grid[bool]
}

func New(x, y int) LCD {
	g := grid.New[bool](x, y)
	return LCD{g: g}
}

func (l *LCD) String() string {
	r := ""
	for i := range l.g.Y() {
		if i != 0 {
			r += "\n"
		}
		for j := range l.g.X() {
			if l.g.At(j, i) {
				r += "#"
			} else {
				r += "."
			}
		}
	}
	return r
}

func (l *LCD) Read() (string, error) {
	var text string
	for i := range l.g.X() / glyphWidth {
		glyph := l.readGlyph(i)
		char, err := glyph.Rune()
		if err != nil {
			return text, err
		}
		text += string(char)
	}
	return text, nil
}

func (l *LCD) readGlyph(index int) Glyph {
	vals := []bool{}
	for y := range l.g.Y() {
		for x := index * glyphWidth; x < ((index + 1) * glyphWidth); x++ {
			vals = append(vals, l.g.At(x, y))
		}
	}
	return Glyph(grid.NewFromSlice(glyphWidth, glyphHeight, vals))
}

func (l *LCD) Do(op string) error {
	split := strings.Split(op, " ")
	if split[0] == "rect" {
		rect := strings.Split(split[1], "x")

		x, err := strconv.Atoi(rect[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(rect[1])
		if err != nil {
			return err
		}
		l.rect(x, y)
	} else {
		n, err := strconv.Atoi(strings.Split(split[2], "=")[1])
		if err != nil {
			return err
		}

		by, err := strconv.Atoi(split[4])
		if err != nil {
			return err
		}

		if split[1] == "column" {
			l.rotCol(n, by)
		} else {
			l.rotRow(n, by)
		}
	}

	return nil
}

func (l *LCD) rect(x, y int) {
	for i := range x {
		for j := range y {
			l.g.Set(i, j, true)
		}
	}
}

func (l *LCD) rotCol(n, by int) {
	for range by {
		next := l.g.At(n, l.g.Y()-1)
		for i := range l.g.Y() {
			tmp := l.g.At(n, i)
			l.g.Set(n, i, next)
			next = tmp
		}
	}
}

func (l *LCD) rotRow(n, by int) {
	for range by {
		next := l.g.At(l.g.X()-1, n)
		for i := range l.g.X() {
			tmp := l.g.At(i, n)
			l.g.Set(i, n, next)
			next = tmp
		}
	}
}

func (l *LCD) CountLit() int {
	r := 0
	for i := range l.g.X() {
		for j := range l.g.Y() {
			if l.g.At(i, j) {
				r += 1
			}
		}
	}
	return r
}

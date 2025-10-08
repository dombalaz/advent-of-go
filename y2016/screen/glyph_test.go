package screen_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016/screen"
)

func TestNewGlyph(t *testing.T) {
	g := screen.NewGlyph()

	actual := g.String()

	w := `.....
.....
.....
.....
.....
.....`
	if actual != w {
		t.Errorf("g.String() = %v, want %v", actual, w)
	}
}

func TestAvailableAlphabet(t *testing.T) {
	alphabet := "BCEFJLOPUZ"

	for _, r := range alphabet {
		glyph, err := screen.NewGlyphFromRune(r)
		if err != nil {
			t.Errorf("screen.NewGlyphFromRune(%v) = %v, want nil", r, err)
		}
		actual, err := glyph.Rune()
		if err != nil {
			t.Errorf("glyph.Rune() = %v, want nil", err)
		}
		if r != actual {
			t.Errorf("glyph.Rune() = %v, want %v", actual, r)
		}
	}
}

func TestNewGlyphFromRune(t *testing.T) {
	r := 'x'
	if _, err := screen.NewGlyphFromRune(r); err == nil {
		t.Errorf("screen.NewGlyphFromRune(%v) = nil, want err", r)
	}
}

func TestGlyphRune(t *testing.T) {
	glyph := screen.NewGlyph()
	if _, err := glyph.Rune(); err == nil {
		t.Errorf("glyph.Rune() = nil, want err")
	}
}

package screen_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016/screen"
)

func TestLcdNew(t *testing.T) {
	lcd := screen.New(3, 3)
	actual := lcd.String()
	w := `...
...
...`

	if actual != w {
		t.Errorf("lcd.String() = %v, want %v", actual, w)
	}
}

func TestLcdScreenOperation(t *testing.T) {
	lcd := screen.New(7, 3)
	wLights := 6

	tests := []struct {
		op string
		w  string
	}{
		{
			op: "rect 3x2",
			w: `###....
###....
.......`,
		},
		{
			op: "rotate column x=1 by 1",
			w: `#.#....
###....
.#.....`,
		},
		{
			op: "rotate row y=0 by 4",
			w: `....#.#
###....
.#.....`,
		},
		{
			op: "rotate column x=1 by 1",
			w: `.#..#.#
#.#....
.#.....`,
		},
	}

	for _, test := range tests {
		t.Run(test.op, func(t *testing.T) {
			lcd.Do(test.op)

			cLights := lcd.CountLit()
			if cLights != wLights {
				t.Errorf("lcd.CountLit() = %v, want %v", cLights, wLights)
			}

			actual := lcd.String()
			if actual != test.w {
				t.Errorf("lcd.String() = %v, want %v", actual, test.w)
			}
		})
	}
}

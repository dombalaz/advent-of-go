package day10_test

import (
	"strings"
	"testing"

	"github.com/dombalaz/advent-of-go/y2016/day10"
)

func TestMicrochipBots(t *testing.T) {
	instructions := `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2
`
	set, err := day10.NewInstructions(strings.NewReader(instructions))
	if err != nil {
		t.Errorf("failed to create instructions: %v", err)
	}
	set.Run()

	tests := []struct {
		a, b int
		w    int
	}{
		{a: 3, b: 5, w: 0},
		{a: 2, b: 3, w: 1},
		{a: 2, b: 5, w: 2},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			v, ok := set.BotIdWithValues(test.a, test.b)
			if !ok {
				t.Errorf("bot not found for values %v, %v, want: %v", test.a, test.b, test.w)
			}
			if v != test.w {
				t.Errorf("found bot %v for values %v, %v, want: %v", v, test.a, test.b, test.w)
			}
		})
	}

	w := 30
	v := set.Multiply()
	if v != w {
		t.Errorf("set.Multiply() = %v, want %v", v, w)
	}
}

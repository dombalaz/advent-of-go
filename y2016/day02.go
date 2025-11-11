package y2016

import (
	"bufio"
	"context"
	"io"

	"github.com/dombalaz/advent-of-go/grid"
)

type Keypad struct {
	g      grid.Grid[rune]
	cx, cy int
}

func NewKeypad() Keypad {
	m := grid.New[rune](3, 3)
	m.Fill([]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'})
	return Keypad{
		g:  m,
		cx: 1,
		cy: 1,
	}
}

func NewHexKeypad() Keypad {
	m := grid.New[rune](5, 5)
	m.Fill([]rune{' ', ' ', '1', ' ', ' ', ' ', '2', '3', '4', ' ', '5', '6', '7', '8', '9', ' ', 'A', 'B', 'C', ' ', ' ', ' ', 'D', ' ', ' '})
	return Keypad{
		g:  m,
		cx: 0,
		cy: 2,
	}
}

func (kp *Keypad) move(i rune) {
	switch i {
	case 'U':
		kp.moveUp()
	case 'R':
		kp.moveRight()
	case 'D':
		kp.moveDown()
	case 'L':
		kp.moveLeft()
	}
}

func (kp *Keypad) moveUp() {
	if kp.cy == 0 || kp.g.At(kp.cx, kp.cy-1) == ' ' {
		return
	}
	kp.cy -= 1
}

func (kp *Keypad) moveLeft() {
	if kp.cx == 0 || kp.g.At(kp.cx-1, kp.cy) == ' ' {
		return
	}
	kp.cx -= 1
}

func (kp *Keypad) moveDown() {
	if kp.cy+1 == kp.g.Y() || kp.g.At(kp.cx, kp.cy+1) == ' ' {
		return
	}
	kp.cy += 1
}

func (kp *Keypad) moveRight() {
	if kp.cx+1 == kp.g.X() || kp.g.At(kp.cx+1, kp.cy) == ' ' {
		return
	}
	kp.cx += 1
}

type Solver02 struct{}

func (s *Solver02) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	return s.solve(ctx, r, NewKeypad())
}

func (s *Solver02) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	return s.solve(ctx, r, NewHexKeypad())
}

func (s *Solver02) solve(ctx context.Context, r io.Reader, kp Keypad) (string, error) {
	ch := make(chan string)
	var err error
	go func() {
		err = scan(r, ch, bufio.ScanLines)
	}()
	var result string
	for line := range ch {
		for _, v := range line {
			kp.move(v)
		}
		result += string(kp.g.At(kp.cx, kp.cy))
	}
	return result, err
}

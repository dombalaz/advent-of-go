package y2016

import (
	"strings"
)

type Matrix[T any] struct {
	x, y int
	data []T
}

func NewMatrix[T any](x, y int) Matrix[T] {
	return Matrix[T]{x, y, make([]T, x*y)}
}

func (m *Matrix[T]) Fill(vs []T) {
	for i, v := range vs {
		m.data[i] = v
	}
}

func (m *Matrix[T]) At(x, y int) T {
	return m.data[y*m.x+x]
}

type Keypad struct {
	m      Matrix[rune]
	cx, cy int
}

func NewKeypad() Keypad {
	m := NewMatrix[rune](3, 3)
	m.Fill([]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'})
	return Keypad{
		m:  m,
		cx: 1,
		cy: 1,
	}
}

func NewHexKeypad() Keypad {
	m := NewMatrix[rune](5, 5)
	m.Fill([]rune{' ', ' ', '1', ' ', ' ', ' ', '2', '3', '4', ' ', '5', '6', '7', '8', '9', ' ', 'A', 'B', 'C', ' ', ' ', ' ', 'D', ' ', ' '})
	return Keypad{
		m:  m,
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
	if kp.cy == 0 || kp.m.At(kp.cx, kp.cy-1) == ' ' {
		return
	}
	kp.cy -= 1
}

func (kp *Keypad) moveLeft() {
	if kp.cx == 0 || kp.m.At(kp.cx-1, kp.cy) == ' ' {
		return
	}
	kp.cx -= 1
}

func (kp *Keypad) moveDown() {
	if kp.cy+1 == kp.m.y || kp.m.At(kp.cx, kp.cy+1) == ' ' {
		return
	}
	kp.cy += 1
}

func (kp *Keypad) moveRight() {
	if kp.cx+1 == kp.m.x || kp.m.At(kp.cx+1, kp.cy) == ' ' {
		return
	}
	kp.cx += 1
}

func prepareInput2(in string) []string {
	return strings.Split(strings.TrimSpace(in), "\n")
}

func Solve2(in string) (string, string) {
	directions := prepareInput2(in)
	var r1, r2 string
	kp := NewKeypad()
	for _, line := range directions {
		for _, v := range line {
			kp.move(v)
		}
		r1 += string(kp.m.At(kp.cx, kp.cy))
	}
	kp = NewHexKeypad()
	for _, line := range directions {
		for _, v := range line {
			kp.move(v)
		}
		r2 += string(kp.m.At(kp.cx, kp.cy))
	}
	return r1, r2
}

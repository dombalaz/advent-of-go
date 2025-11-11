package y2016

import (
	"context"
	"io"
	"strconv"
)

type Direction int

const (
	// Don't change the order.
	directionUp = iota
	directionRight
	directionDown
	directionLeft
)

func Abs(x int) int {
	mask := x >> 31
	return (x ^ mask) - mask
}

type location struct {
	x int
	y int
}

type gridWalk struct {
	l                   location
	d                   Direction
	visited             map[location]struct{}
	secondVisitDistance int
}

func (gw *gridWalk) turn(where byte) {
	if where == 'R' {
		gw.d += 1
	} else {
		gw.d += directionLeft
	}
	gw.d %= directionLeft + 1
}

func (gw *gridWalk) walk(c int) {
	var walkF func(gw *gridWalk)
	switch gw.d {
	case directionUp:
		walkF = func(gw *gridWalk) {
			gw.l.y += 1
		}
	case directionRight:
		walkF = func(gw *gridWalk) {
			gw.l.x += 1
		}
	case directionDown:
		walkF = func(gw *gridWalk) {
			gw.l.y -= 1
		}
	case directionLeft:
		walkF = func(gw *gridWalk) {
			gw.l.x -= 1
		}
	}
	for range c {
		walkF(gw)
		gw.checkPosition()
	}
}

func (gw *gridWalk) distance() int {
	return Abs(gw.l.x) + Abs(gw.l.y)
}

func (gw *gridWalk) saveLocation() {
	gw.visited[location{x: gw.l.x, y: gw.l.y}] = struct{}{}
}

func (gw *gridWalk) checkPosition() {
	if gw.secondVisitDistance != 0 {
		return
	}
	if _, ok := gw.visited[location{x: gw.l.x, y: gw.l.y}]; ok {
		gw.secondVisitDistance = gw.distance()
	} else {
		gw.saveLocation()
	}
}

func newGridWalk() gridWalk {
	return gridWalk{
		l:                   location{x: 0, y: 0},
		d:                   directionUp,
		visited:             make(map[location]struct{}),
		secondVisitDistance: 0,
	}
}

type Solver01 struct{}

func (s *Solver01) solve(ch <-chan string) (string, string) {
	gw := newGridWalk()

	for v := range ch {
		c, _ := strconv.Atoi(v[1:])
		gw.turn(v[0])
		gw.walk(c)
	}

	r1 := strconv.FormatInt(int64(gw.distance()), 10)
	r2 := strconv.FormatInt(int64(gw.secondVisitDistance), 10)
	return r1, r2
}

func (s *Solver01) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	ch := make(chan string)
	var err error
	go func() {
		err = scan(r, ch, scanCSV)
	}()
	result, _ := s.solve(ch)
	return result, err
}

func (s *Solver01) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	ch := make(chan string)
	var err error
	go func() {
		err = scan(r, ch, scanCSV)
	}()
	_, result := s.solve(ch)
	return result, err
}

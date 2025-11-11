package core

import (
	"context"
	"io"

	"github.com/dombalaz/advent-of-go/y2016"
)

type Day int

type ProblemSolver interface {
	SolveP1(ctx context.Context, r io.Reader) (string, error)
	SolveP2(ctx context.Context, r io.Reader) (string, error)
}

func NewProblemSolver(day Day) ProblemSolver {
	switch day {
	case Day01:
		return &y2016.Solver01{}
	case Day02:
		return &y2016.Solver02{}
	case Day03:
		return &y2016.Solver03{}
	case Day04:
		return &y2016.Solver04{}
	case Day05:
		return &y2016.Solver05{}
	case Day06:
		return &y2016.Solver06{}
	case Day07:
		return &y2016.Solver07{}
	case Day08:
		return &y2016.Solver08{}
	case Day09:
		return &y2016.Solver09{}
	}
	return nil
}

const (
	Day01 = iota + 1
	Day02
	Day03
	Day04
	Day05
	Day06
	Day07
	Day08
	Day09
)

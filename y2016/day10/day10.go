package day10

import (
	"context"
	"errors"
	"io"
	"strconv"
)

const w1 = 17
const w2 = 61

type Solver struct{}

func (s *Solver) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	is, err := NewInstructions(r)
	if err != nil {
		return "", err
	}
	is.Run()

	v, ok := is.BotIdWithValues(w1, w2)
	if !ok {
		return "", errors.New("bot not found")
	}
	return strconv.FormatInt(int64(v), 10), nil
}

func (s *Solver) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	is, err := NewInstructions(r)
	if err != nil {
		return "", err
	}
	is.Run()
	return strconv.FormatInt(int64(is.Multiply()), 10), nil
}

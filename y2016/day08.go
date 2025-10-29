package y2016

import (
	"bufio"
	"context"
	"io"
	"strconv"

	"github.com/dombalaz/advent-of-go/y2016/screen"
)

type Solver08 struct{}

func (s *Solver08) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	f := func(lcd *screen.LCD) (string, error) {
		return strconv.FormatInt(int64(lcd.CountLit()), 10), nil
	}

	return s.solve(ctx, r, f)
}

func (s *Solver08) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	f := func(lcd *screen.LCD) (string, error) {
		return lcd.Read()
	}

	return s.solve(ctx, r, f)
}

func (s *Solver08) solve(ctx context.Context, r io.Reader, f func(lcd *screen.LCD) (string, error)) (string, error) {
	ch := make(chan string)
	go scan(r, ch, bufio.ScanLines)

	run := func(ch <-chan string) (screen.LCD, error) {
		lcd := screen.New(50, 6)
		for v := range ch {
			if err := lcd.Do(v); err != nil {
				return lcd, nil
			}
		}
		return lcd, nil
	}

	lcd, err := run(ch)
	if err != nil {
		return "", err
	}

	return f(&lcd)
}

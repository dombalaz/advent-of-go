package y2016

import (
	"strconv"
	"strings"

	"github.com/dombalaz/advent-of-go/y2016/screen"
)

type Solver08 struct{}

func (s *Solver08) SolveP1(in string) string {
	lines := strings.Split(in, "\n")
	lcd := screen.New(50, 6)
	for _, v := range lines {
		if err := lcd.Do(v); err != nil {
			return ""
		}
	}
	return strconv.FormatInt(int64(lcd.CountLit()), 10)
}

func (s *Solver08) SolveP2(in string) string {
	lines := strings.Split(in, "\n")
	lcd := screen.New(50, 6)
	for _, v := range lines {
		if err := lcd.Do(v); err != nil {
			return ""
		}
	}
	str, _ := lcd.Read()
	return str
}

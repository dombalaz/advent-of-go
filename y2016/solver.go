package y2016

type Day int

type Solver interface {
	SolveP1(in string) string
	SolveP2(in string) string
}

const (
	Day01 = iota + 1
	Day02
	Day03
	Day04
	Day05
	Day06
)

func CreateSolver(day Day) Solver {
	switch day {
	case Day01:
		return &Solver01{}
	case Day02:
		return &Solver02{}
	case Day03:
		return &Solver03{}
	case Day04:
		return &Solver04{}
	case Day05:
		return &Solver05{}
	case Day06:
		return &Solver06{}
	}
	return nil
}

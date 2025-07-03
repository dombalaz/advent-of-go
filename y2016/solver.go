package y2016

type Day int

type Solver interface {
	SolveP1(in string) string
	SolveP2(in string) string
}

const (
	Day01 = iota + 1
	Day02
)

func CreateSolver(day Day) Solver {
	switch day {
	case Day01:
		return &Solver01{}
	case Day02:
		return &Solver02{}
	}
	return nil
}

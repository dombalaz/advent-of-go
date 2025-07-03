package y2016

type Day int

type Solver func(in string) (string, string)

const (
	Day01 = iota + 1
	Day02
)

func CreateSolver(day Day) Solver {
	switch day {
	case Day01:
		return Solve
	case Day02:
		return Solve2
	}
	return nil
}

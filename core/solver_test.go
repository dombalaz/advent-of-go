package core_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/core"
)

func TestInvalidDay(t *testing.T) {
	solver := core.NewProblemSolver(0)
	if solver != nil {
		t.Errorf("y2016.CreateSolver(0) = %v, want nil", solver)
	}
}

func TestNewProblemSolver(t *testing.T) {
	implementedDays := 8
	for i := 1; i <= implementedDays; i++ {
		solver := core.NewProblemSolver(core.Day(i))
		if solver == nil {
			t.Errorf("y2016.CreateSolver(0) = nil, want object")
		}
	}
}

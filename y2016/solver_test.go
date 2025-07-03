package y2016_test

import (
	"testing"

	"github.com/dombalaz/advent-of-go/y2016"
)

func TestInvalidDay(t *testing.T) {
	solver := y2016.CreateSolver(0)
	if solver != nil {
		t.Errorf("y2016.CreateSolver(0) = %v, want nil", solver)
	}
}

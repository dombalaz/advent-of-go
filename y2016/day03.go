package y2016

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidTriangle(a, b, c int) bool {
	if a < b+c && b < a+c && c < a+b {
		return true
	}
	return false
}

type Solver03 struct{}

func countTriangles(in []string) int {
	valid := 0
	for _, v := range in {
		sides := strings.Fields(v)

		a, _ := strconv.Atoi(sides[0])
		b, _ := strconv.Atoi(sides[1])
		c, _ := strconv.Atoi(sides[2])

		if isValidTriangle(a, b, c) {
			valid += 1
		}
	}
	return valid
}

func (s *Solver03) SolveP1(in string) string {
	triangles := strings.Split(in, "\n")
	return strconv.FormatInt(int64(countTriangles(triangles)), 10)
}

func (s *Solver03) SolveP2(in string) string {
	triangles := strings.Split(in, "\n")

	var valid int
	var converted [3]string
	for i := 0; i < len(triangles); {
		sides1 := strings.Fields(triangles[i])
		sides2 := strings.Fields(triangles[i+1])
		sides3 := strings.Fields(triangles[i+2])

		converted[0] = fmt.Sprintf("%v %v %v", sides1[0], sides2[0], sides3[0])
		converted[1] = fmt.Sprintf("%v %v %v", sides1[1], sides2[1], sides3[1])
		converted[2] = fmt.Sprintf("%v %v %v", sides1[2], sides2[2], sides3[2])

		valid += countTriangles(converted[:])

		i += 3
	}

	return strconv.FormatInt(int64(valid), 10)
}

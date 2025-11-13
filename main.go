package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dombalaz/advent-of-go/core"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid number of arguments")
		os.Exit(1)
	}
	day, _ := strconv.Atoi(os.Args[1])
	filename := os.Args[2]
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("failed to read from file: %v", err)
		os.Exit(1)
	}

	s := strings.TrimSpace(string(bytes))

	solver := core.NewProblemSolver(core.Day(day))
	if solver == nil {
		fmt.Printf("failed to create a solver for day: %v", day)
		os.Exit(1)
	}

	ctx := context.Background()

	r1, err := solver.SolveP1(ctx, strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	r2, err := solver.SolveP2(ctx, strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("========== p1 ==========")
	fmt.Println(r1)
	fmt.Println("========== p2 ==========")
	fmt.Println(r2)
	fmt.Println("========================")
}

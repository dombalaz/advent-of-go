package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dombalaz/advent-of-go/y2016"
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

	solver := y2016.CreateSolver(y2016.Day(day))
	if solver == nil {
		fmt.Printf("failed to create a solver for day: %v", day)
		os.Exit(1)
	}

	r1 := solver.SolveP1(s)
	r2 := solver.SolveP2(s)
	fmt.Println("========== p1 ==========")
	fmt.Println(r1)
	fmt.Println("========== p2 ==========")
	fmt.Println(r2)
	fmt.Println("========================")
}

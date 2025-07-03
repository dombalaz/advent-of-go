package main

import (
	"fmt"
	"os"

	"github.com/dombalaz/advent-of-go/y2016"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid number of arguments")
		os.Exit(1)
	}
	filename := os.Args[1]
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("failed to read from file: %v", err)
		os.Exit(1)
	}

	s := string(bytes)

	r1, r2 := y2016.Solve(s)
	fmt.Printf("result: %v, %v", r1, r2)
}

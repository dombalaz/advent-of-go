package y2016

import (
	"bufio"
	"context"
	"io"
	"strconv"
	"strings"
)

type Solver07 struct{}

const minABASize = 3 + 3 + 2 // aba[bab]
const minABBASize = 4        // abba

func isABBA(s string) bool {
	return s[0] == s[3] && s[1] == s[2]
}

func supportsABBA(s string) bool {
	if len(s) < minABBASize {
		return false
	}

	// algorithm checks chars forward so stop 3 chars before end
	stop := len(s) - 3
	var is bool
	var inside bool
	for i := range stop {
		if s[i] == s[i+1] {
			// characters has to be different in pair
			continue
		}
		if s[i] == '[' || s[i] == ']' {
			inside = !inside
			continue
		}
		if isABBA(s[i : i+4]) {
			if inside {
				// if the sequence is inside the brackets we return false
				return false
			}
			is = true
		}
	}
	return is
}

// tag is either ABA or BAB
func isTag(s string) bool {
	if strings.ContainsAny(s, "[]") {
		return false
	}
	if s[0] == s[2] && s[0] != s[1] {
		return true
	}
	return false
}

// expects already that it is a tag
func isABABAB(s1 string, s2 string) bool {
	if s1[0] == s2[1] && s1[1] == s2[0] {
		return true
	}
	return false
}

func supportsSSL(s string) bool {
	var abas []string
	var babs []string

	if len(s) < minABASize {
		return false
	}

	stop := len(s) - 2
	var inside bool
	for i := range stop {
		if s[i] == '[' || s[i] == ']' {
			inside = !inside
			continue
		}
		sub := s[i : i+3]
		if isTag(sub) {
			if inside {
				babs = append(babs, sub)
			} else {
				abas = append(abas, sub)
			}
		}
	}
	for _, v1 := range abas {
		for _, v2 := range babs {
			if isABABAB(v1, v2) {
				return true
			}
		}
	}
	return false
}

func (s *Solver07) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	ch := make(chan string)
	go scan(r, ch, bufio.ScanLines)

	run := func(ch <-chan string) int64 {
		var c int64
		for v := range ch {
			if supportsABBA(v) {
				c++
			}
		}
		return c
	}

	return strconv.FormatInt(run(ch), 10), nil
}

func (s *Solver07) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	ch := make(chan string)
	go scan(r, ch, bufio.ScanLines)

	run := func(ch <-chan string) int64 {
		var c int64
		for v := range ch {
			if supportsSSL(v) {
				c++
			}
		}
		return c
	}

	return strconv.FormatInt(run(ch), 10), nil
}

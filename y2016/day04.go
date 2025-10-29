package y2016

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/dombalaz/advent-of-go/cipher/caesar"
)

type Solver04 struct{}

func (s *Solver04) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	ch := make(chan string)
	process := func(ch <-chan string) int64 {
		var sum int
		for v := range ch {
			r := parseRoom(v)
			var chars []rune
			str := r.letters
			for range 5 {
				char := mostFreqRune(str)
				str = removeRune(str, char)
				chars = append(chars, char)
			}
			if r.checksum == string(chars) {
				sum += r.id
			}
		}
		return int64(sum)
	}

	go scan(r, ch, bufio.ScanLines)
	return strconv.FormatInt(process(ch), 10), nil
}

func (s *Solver04) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	ch := make(chan string)
	process := func(ch <-chan string) string {
		var str string
		for v := range ch {
			r := parseRoom(v)
			var buf bytes.Buffer
			writer, _ := caesar.NewWriter(&buf, r.id)
			writer.Write([]byte(r.letters))
			if strings.HasPrefix(buf.String(), "north") {
				str = strconv.FormatInt(int64(r.id), 10)
				break
			}
		}
		return str
	}

	go scan(r, ch, bufio.ScanLines)
	return process(ch), nil
}

type room struct {
	id       int
	letters  string
	checksum string
}

func parseRoom(s string) room {
	var r room

	iCheck := strings.Index(s, "[")
	r.checksum = s[iCheck+1 : len(s)-1]

	iId := strings.LastIndex(s, "-")
	r.id, _ = strconv.Atoi(s[iId+1 : iCheck])

	r.letters = removeRune(s[:iId], '-')

	return r
}

// Fins most frequent chars, if tie return alphabetically.
func mostFreqRune(s string) rune {
	m := make(map[rune]int)
	var max int
	var chars []rune

	for _, v := range s {
		m[v]++
		if max < m[v] {
			max = m[v]
		}
	}

	for k, v := range m {
		if max == v {
			chars = append(chars, k)
		}
	}

	slices.Sort(chars)

	return chars[0]
}

func removeRune(s string, r rune) string {
	return strings.Join(strings.Split(s, string(r)), "")
}

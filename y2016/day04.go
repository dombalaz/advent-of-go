package y2016

import (
	"bytes"
	"slices"
	"strconv"
	"strings"

	"github.com/dombalaz/advent-of-go/cipher/caesar"
)

type Solver04 struct{}

func (s *Solver04) SolveP1(in string) string {
	l := strings.Split(in, "\n")
	var sum int

	for _, v := range l {
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

	return strconv.FormatInt(int64(sum), 10)
}

func (s *Solver04) SolveP2(in string) string {
	l := strings.Split(in, "\n")
	res := ""
	for _, v := range l {
		r := parseRoom(v)
		var buf bytes.Buffer
		writer, _ := caesar.NewWriter(&buf, r.id)
		writer.Write([]byte(r.letters))
		if strings.HasPrefix(buf.String(), "north") {
			res = strconv.FormatInt(int64(r.id), 10)
			break
		}
	}
	return res
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

package y2016

import (
	"slices"
	"strconv"
	"strings"
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
	var sum int
	res := ""
	for _, v := range l {
		r := parseRoom(v)
		var chars []rune
		str := r.letters
		for range 5 {
			char := mostFreqRune(r.letters)
			str = removeRune(str, char)
			chars = append(chars, char)
		}
		if r.checksum == string(chars) {
			sum += r.id
		}
		str = string(CaesarCipherShift([]byte(r.letters), r.id))
		if strings.HasPrefix(str, "north") {
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

func CaesarCipherShift(s []byte, c int) []byte {
	c = c % 26
	var r byte
	for i := range s {
		if s[i] < 'a' || 'z' < s[i] {
			continue
		}

		r = byte(int(s[i]) + c)
		if 'z' < r {
			r = byte(int(r) - 'z' + 'a' - 1)
		}

		s[i] = r
	}
	return s
}

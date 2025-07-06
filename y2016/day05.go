package y2016

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

type Solver05 struct{}

func (s *Solver05) SolveP1(in string) string {
	r := ""
	for i := int64(0); ; i++ {
		str := in + strconv.FormatInt(i, 10)
		sum := md5.Sum([]byte(str))
		str = hex.EncodeToString(sum[:])
		if strings.HasPrefix(str, "00000") {
			r += string(str[5])
		}
		if len(r) == 8 {
			return r
		}
	}
}

func (s *Solver05) SolveP2(in string) string {
	r := make([]rune, 8)
	for i := int64(0); ; i++ {
		str := in + strconv.FormatInt(i, 10)
		sum := md5.Sum([]byte(str))
		str = hex.EncodeToString(sum[:])
		if !strings.HasPrefix(str, "00000") {
			continue
		}
		if str[5] < '0' || '7' < str[5] {
			continue
		}
		pos := str[5] - '0'
		if r[pos] != 0 {
			continue
		}
		r[pos] = rune(str[6])
		if strings.IndexRune(string(r), 0) == -1 {
			return string(r)
		}
	}
}

package y2016

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"strings"
)

type Solver05 struct{}

func (s *Solver05) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", nil
	}

	run := func(str string) string {
		var result string
		for i := int64(0); ; i++ {
			str := str + strconv.FormatInt(i, 10)
			sum := md5.Sum([]byte(str))
			str = hex.EncodeToString(sum[:])
			if strings.HasPrefix(str, "00000") {
				result += string(str[5])
			}
			if len(result) == 8 {
				return result
			}
		}
	}

	return run(string(b)), nil
}

func (s *Solver05) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", nil
	}

	run := func(str string) string {
		result := make([]rune, 8)
		for i := int64(0); ; i++ {
			str := str + strconv.FormatInt(i, 10)
			sum := md5.Sum([]byte(str))
			str = hex.EncodeToString(sum[:])
			if !strings.HasPrefix(str, "00000") {
				continue
			}
			if str[5] < '0' || '7' < str[5] {
				continue
			}
			pos := str[5] - '0'
			if result[pos] != 0 {
				continue
			}
			result[pos] = rune(str[6])
			if !strings.ContainsRune(string(result), 0) {
				return string(result)
			}
		}
	}

	return run(string(b)), nil
}

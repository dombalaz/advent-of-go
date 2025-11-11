package y2016

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Solver09 struct{}

func parseMarker(s string) (int, int, error) {
	split := strings.Split(s, "x")
	if len(split) != 2 {
		return 0, 0, fmt.Errorf("invalid marker: %s", s)
	}
	c, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid marker: %s", s)
	}
	r, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid marker: %s", s)
	}
	return c, r, nil
}

func DetermineSize(r io.Reader, v2 bool) (int, error) {
	var size int
	br := bufio.NewReader(r)

	for {
		// Read until mark
		data, err := br.ReadBytes('(')
		if err != nil {
			// If error or EOF just return what was read
			if err == io.EOF {
				size += len(data)
				err = nil
			}
			return size, err
		}

		size += len(data) - 1

		// Process block
		data, err = br.ReadBytes(')')
		if err != nil {
			return size, errors.New("invalid block")
		}

		trimmed := string(data[:len(data)-1])
		count, repeat, err := parseMarker(trimmed)
		if err != nil {
			return size, err
		}

		var times int
		if v2 {
			// Read whole compressed part
			compressed := make([]byte, count)
			_, err = io.ReadFull(br, compressed)
			if err != nil {
				return size, fmt.Errorf("compressed part read: %v", err)
			}
			// Determine size for the compressed part
			compressedSize, err := DetermineSize(bytes.NewReader(compressed), true)
			if err != nil {
				return size, err
			}
			times = compressedSize
		} else {
			_, err = br.Discard(count)
			if err != nil {
				// Happens when not all were discarded so invalid block
				return size, errors.New("invalid block")
			}
			times = count
		}

		size += repeat * times
	}
}

func (s *Solver09) SolveP1(ctx context.Context, r io.Reader) (string, error) {
	size, err := DetermineSize(r, false)
	return strconv.FormatInt(int64(size), 10), err
}

func (s *Solver09) SolveP2(ctx context.Context, r io.Reader) (string, error) {
	size, err := DetermineSize(r, true)
	return strconv.FormatInt(int64(size), 10), err
}

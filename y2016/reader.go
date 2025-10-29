package y2016

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

func scan(r io.Reader, ch chan<- string, split bufio.SplitFunc) error {
	defer close(ch)
	scanner := bufio.NewScanner(r)
	scanner.Split(split)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func scanCSV(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if i := bytes.IndexByte(data, ','); i >= 0 {
		field := strings.TrimSpace(string(data[:i]))
		return i + 1, []byte(field), nil
	}
	if atEOF && len(data) > 0 {
		field := strings.TrimSpace(string(data))
		return len(data), []byte(field), nil
	}
	return 0, nil, nil
}

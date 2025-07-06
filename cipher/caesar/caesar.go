package caesar

import (
	"io"
)

type header struct {
	shift int
}

// Decrypts the data using caesar cipher
type Reader struct {
	header
	r io.Reader
}

func shift(b byte, c int) byte {
	shiftWidth := 26
	switch {
	case 'A' <= b && b <= 'Z':
		return 'A' + byte((((int(b-'A')+c)%shiftWidth)+shiftWidth)%shiftWidth)
	case 'a' <= b && b <= 'z':
		return 'a' + byte((((int(b-'a')+c)%shiftWidth)+shiftWidth)%shiftWidth)
	default:
		return b
	}
}

func (r *Reader) Read(p []byte) (int, error) {
	tmp := make([]byte, len(p))
	n, err := r.r.Read(tmp)
	for i := range n {
		p[i] = shift(tmp[i], -r.shift)
	}
	return n, err
}

func NewReader(r io.Reader, shift int) (*Reader, error) {
	return &Reader{r: r, header: header{shift: shift}}, nil
}

// Encrypts the data using caesar cipher
type Writer struct {
	header
	r io.Writer
}

func (r *Writer) Write(p []byte) (int, error) {
	tmp := make([]byte, len(p))
	for i := range len(p) {
		tmp[i] = shift(p[i], r.shift)
	}
	n, err := r.r.Write(tmp)
	return n, err
}

func NewWriter(r io.Writer, shift int) (*Writer, error) {
	return &Writer{r: r, header: header{shift: shift}}, nil
}

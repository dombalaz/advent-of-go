package caesar_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dombalaz/advent-of-go/cipher/caesar"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func TestReaderShift(t *testing.T) {
	text := "bcdefghijklmnopqrstuvwxyza"
	sr := strings.NewReader(text)
	cr, _ := caesar.NewReader(sr, 1)
	bytes := make([]byte, len(text))
	n, err := cr.Read(bytes)
	if n != len(bytes) || err != nil || string(bytes) != alphabet {
		t.Errorf("caesar reader with shift %v = %v, want %v", 1, string(bytes), alphabet)
	}
}

func TestReaderUpper(t *testing.T) {
	text := strings.ToUpper("bcdefghijklmnopqrstuvwxyza")
	bytes := make([]byte, len(text))
	sr := strings.NewReader(text)
	cr, _ := caesar.NewReader(sr, 1)
	n, err := cr.Read(bytes)
	w := strings.ToUpper(alphabet)
	if n != len(bytes) || err != nil || string(bytes) != w {
		t.Errorf("caesar reader with shift %v = %v, want %v", 1, string(bytes), w)
	}
}

func TestReaderWrap(t *testing.T) {
	bytes := make([]byte, len(alphabet))
	shift := 26
	sr := strings.NewReader(alphabet)
	cr, _ := caesar.NewReader(sr, shift)
	n, err := cr.Read(bytes)
	if n != len(bytes) || err != nil || string(bytes) != alphabet {
		t.Errorf("caesar reader with shift %v = %v, want %v", shift, string(bytes), alphabet)
	}
}

func TestWriterShift(t *testing.T) {
	var buf bytes.Buffer
	writer, _ := caesar.NewWriter(&buf, 1)
	n, err := writer.Write([]byte(alphabet))
	w := "bcdefghijklmnopqrstuvwxyza"
	if n != len(w) || err != nil || buf.String() != w {
		t.Errorf("caesar reader with shift %v = %v, want %v", 1, string(buf.String()), w)
	}
}

func TestWriterNonLetters(t *testing.T) {
	var buf bytes.Buffer
	writer, _ := caesar.NewWriter(&buf, 1)
	n, err := writer.Write([]byte("ab.cd-ef!"))
	w := "bc.de-fg!"
	if n != len(w) || err != nil || buf.String() != w {
		t.Errorf("caesar reader with shift %v = %v, want %v", 1, string(buf.String()), w)
	}
}

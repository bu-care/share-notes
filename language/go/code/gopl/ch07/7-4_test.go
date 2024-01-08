package ch07

import (
	"fmt"
	"io"
	"os"
	"testing"

	"golang.org/x/net/html"
)

type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

// Read implements the io.Reader interface.
func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func NewReader(s string) io.Reader {
	return &Reader{s, 0, -1}
}

func TestNewReader(t *testing.T) {
	res, err := html.Parse(NewReader("<h1>Hello</h1>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse err: %v", err)
		os.Exit(1)
	}
	fmt.Println("res:", res)
}

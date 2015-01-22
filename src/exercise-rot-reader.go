package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	l, err := r13.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i, v := range b {
		if v == 0 {
			break
		}
		var s, e byte = 'A', 'Z'
		if v < 'z' {
			s, e = 'a', 'z'
		}
		if v+13 > e {
			b[i] = (s - 1) + ((v + 13) - e)
		} else {
			b[i] += 13
		}
	}
	return l, io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
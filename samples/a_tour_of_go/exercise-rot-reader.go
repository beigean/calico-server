package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	if err == nil {
		for i, v := range b {
			switch {
			case v >= 'A' && v <= 'Z':
				b[i] = (v-'A'+13)%26 + 'A'
			case v >= 'a' && v <= 'z':
				b[i] = (v-'a'+13)%26 + 'a'
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

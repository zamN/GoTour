package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	R io.Reader
}

func (rot rot13Reader) Read(p []byte) (bytesRead int, err error) {
	bytesRead, err = rot.R.Read(p)
	for i := 0; i < len(p); i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = ((p[i] + 13) % 'z')
			if p[i] < 'a' {
				p[i] = p[i] + 'a' - 1
			}
		} else if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = ((p[i] + 13) % 'Z')
			if p[i] < 'A' {
				p[i] = p[i] + 'A' - 1
			}
		}
	}
	return bytesRead, err
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

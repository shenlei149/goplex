package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	lr := LimitReader(os.Stdin, 20)
	p := make([]byte, 7)
	n, err := lr.Read(p)
	fmt.Println(n, err, string(p[:n]))
	n, err = lr.Read(p)
	fmt.Println(n, err, string(p[:n]))
	n, err = lr.Read(p)
	fmt.Println(n, err, string(p[:n]))
	n, err = lr.Read(p)
	fmt.Println(n, err, string(p[:n]))
}

type LimitedReader struct {
	R io.Reader
	N int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}

	n, err = l.R.Read(p)
	l.N -= int64(n)

	return
}

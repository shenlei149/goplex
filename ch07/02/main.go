package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	w, n := CountingWriter(os.Stdout)
	w.Write([]byte("Hello, world!"))
	w.Write([]byte("func main()"))
	fmt.Println(*n)
}

type CounterWriter struct {
	writer io.Writer
	count  int64
}

func (cw *CounterWriter) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	cw.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CounterWriter{w, 0}
	return cw, &cw.count
}

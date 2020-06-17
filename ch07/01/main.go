package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	s1 := "Hello, world!\nTest WC and LC."
	s2 := "func main"

	var wc WordCounter
	c, _ := wc.Write([]byte(s1))
	fmt.Println(c, wc)
	c, _ = wc.Write([]byte(s2))
	fmt.Println(c, wc)

	var lc LineCounter
	c, _ = lc.Write([]byte(s1))
	fmt.Println(c, lc)
	c, _ = lc.Write([]byte(s2))
	fmt.Println(c, lc)
}

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanWords)
	count := 0
	for s.Scan() {
		count++
	}

	*wc += WordCounter(count)
	return count, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanLines)
	count := 0
	for s.Scan() {
		count++
	}

	*lc += LineCounter(count)
	return count, nil
}

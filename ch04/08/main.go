package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var Stdin io.Reader = os.Stdin
var Stdout io.Writer = os.Stdout
var Stderr io.Writer = os.Stderr

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	letters := 0                    // count of letters
	digits := 0                     // count of digits
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		} else if unicode.IsLetter(r) {
			letters++
		} else if unicode.IsDigit(r) {
			digits++
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Fprintf(Stdout, "rune\tcount\n")
	for c, n := range counts {
		fmt.Fprintf(Stdout, "%q\t%d\n", c, n)
	}

	fmt.Fprint(Stdout, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(Stdout, "%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Fprintf(Stdout, "\n%d invalid UTF-8 characters\n", invalid)
	}

	if letters > 0 {
		fmt.Fprintf(Stdout, "\n%d letters\n", letters)
	}

	if digits > 0 {
		fmt.Fprintf(Stdout, "\n%d digits\n", digits)
	}
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}

	fmt.Printf("  %s\n", comma(""))
}

func comma(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	var buf bytes.Buffer
	var sign int
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		sign++
	}

	point := strings.Index(s, ".")
	integerLength := n - sign
	if point >= 0 {
		integerLength = point
	}

	i := integerLength % 3
	if i == 0 && integerLength != 0 {
		i = 3
	}

	i += sign
	buf.WriteString(s[sign:i])

	for ; i < integerLength+sign; i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}

	if point >= 0 {
		buf.WriteString(".")
		i++

		for ; i < n; i += 3 {
			if i+3 < n {
				buf.WriteString(s[i : i+3])
			} else {
				buf.WriteString(s[i:n])
			}

			if i+3 < n {
				buf.WriteString(",")
			}
		}
	}

	return buf.String()
}

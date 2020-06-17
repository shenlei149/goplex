package main

import (
	"bytes"
	"fmt"
	"os"
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
	i := n % 3
	if i == 0 {
		i = 3
	}

	buf.WriteString(s[0:i])

	for ; i < n; i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

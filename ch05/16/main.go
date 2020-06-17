package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(stringsJoin("#t"))
	fmt.Println(stringsJoin("#t", "hello"))
	fmt.Println(stringsJoin("#t", "hello", "world"))
	fmt.Println(stringsJoin("#t", "hello", "world", "!"))
}

func stringsJoin(sep string, a ...string) string {
	return strings.Join(a, sep)
}

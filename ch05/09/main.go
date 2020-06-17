package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("hello $foo! $foo.", f))
}

func f(str string) string {
	return str + "!"
}

func expand(s string, f func(string) string) string {
	ret := f("foo")
	return strings.ReplaceAll(s, "$foo", ret)
}

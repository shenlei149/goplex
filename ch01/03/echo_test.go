package main

import (
	"strings"
	"testing"
)

func echo1(strings []string) string {
	var s, sep string
	for i := 0; i < len(strings); i++ {
		s += sep + strings[i]
		sep = " "
	}
	return s
}

func echo2(strings []string) string {
	s, sep := "", ""
	for _, str := range strings {
		s += sep + str
		sep = " "
	}
	return s
}

func echo3(str []string) string {
	return strings.Join(str, " ")
}

func benchmarkEcho(b *testing.B, echo func([]string) string, args []string) {
	for i := 0; i < b.N; i++ {
		echo(args)
	}
}

var args []string = []string{"Experiment", "to", "measure", "the", "difference", "in",
	"running", "time", "between", "our", "potentially", "inefficient", "versions", "and",
	"the", "one", "that", "uses", "strings.Join.", "(Section", "1.6", "illustrates", "part",
	"of", "the", "time", "package,", "and", "Sec", "tion", "11.4", "shows", "how", "to", "write",
	"benchmark", "tests", "for", "systematic", "performance", "evaluation.)"}

func BenchmarkEcho1(b *testing.B) {
	benchmarkEcho(b, echo1, args)
}

func BenchmarkEcho2(b *testing.B) {
	benchmarkEcho(b, echo2, args)
}

func BenchmarkEcho3(b *testing.B) {
	benchmarkEcho(b, echo3, args)
}

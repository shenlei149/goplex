package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"a b c d e", " ", 5},
		{"", ":", 1},
		{"abc", " ", 1},
	}

	for _, test := range tests {
		s, sep, want := test.s, test.sep, test.want
		words := strings.Split(s, sep)
		if got := len(words); got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", s, sep, got, want)
		}
	}
}

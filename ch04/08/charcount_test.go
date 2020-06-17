package main

import (
	"bytes"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tcs = []struct {
		input   string
		expects string
	}{
		{
			"Hello\xa0",
			"rune\tcount\n" +
			"'H'\t1\n" +
			"'e'\t1\n" +
			"'l'\t2\n" +
			"'o'\t1\n" +
			"\nlen\tcount\n" +
			"1\t5\n" +
			"2\t0\n" +
			"3\t0\n" +
			"4\t0\n" +
			"\n1 invalid UTF-8 characters\n"+
			"\n5 letters\n",
		},
	}

	for _, tc := range tcs {
		Stdin = bytes.NewBufferString(tc.input)
		Stdout = new(bytes.Buffer)
		main()
		ret := Stdout.(*bytes.Buffer).String()

		if ret != tc.expects {
			t.Errorf("Failed count chars.\nInput: %q,\nExpects:\n%q,\nActual:\n%q", tc.input, tc.expects, ret)
		}
	}
}

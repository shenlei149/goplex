package cycle

import "testing"

func TestHasCycle(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}

	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	d, e, f := &link{value: "d"}, &link{value: "e"}, &link{value: "f"}
	d.tail, e.tail, f.tail = e, f, nil

	m := map[string]link{
		"a": *a,
		"b": *b,
		"c": *c,
	}

	testdata := []struct {
		x    interface{}
		want bool
	}{
		{a, true},
		{d, false},
		{1, false},
		{true, false},
		{"abc", false},
		{m, true},
	}

	for _, test := range testdata {
		if HasCycle(test.x) != test.want {
			t.Errorf("HasCycle(%v) = %t", test.x, !test.want)
		}
	}
}

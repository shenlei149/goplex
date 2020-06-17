package main

import (
	"fmt"
	"testing"
)

func checkValues(t *testing.T, set *IntSet, want string) {
	if set.String() != want {
		t.Fatalf("set.String()=%q, want %q", set.String(), want)
	}
}

func TestIntSet(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	checkValues(t, &x, "{1 9 144}")

	y.Add(9)
	y.Add(42)
	checkValues(t, &y, "{9 42}")

	x.UnionWith(&y)
	checkValues(t, &x, "{1 9 42 144}")

	if !x.Has(9) || x.Has(123) {
		t.Fatal("x.Has(9)!=ture or x.Has(123)=false")
	}

	if x.Len() != 4 || y.Len() != 2 {
		t.Fatal("x.Len()!=4 or y.Len()!=2")
	}

	x.Remove(244)
	checkValues(t, &x, "{1 9 42 144}")

	x.Remove(42)
	checkValues(t, &x, "{1 9 144}")

	x.Clear()
	checkValues(t, &x, "{}")

	z := y.Copy()
	y.Add(188)
	checkValues(t, z, "{9 42}")
	checkValues(t, &y, "{9 42 188}")

	z = y.Copy()
	checkValues(t, z, "{9 42 188}")

	z.AddAll()
	checkValues(t, z, "{9 42 188}")

	z.AddAll(1, 68, 138)
	checkValues(t, z, "{1 9 42 68 138 188}")

	x.AddAll(9, 42, 49)
	checkValues(t, &x, "{9 42 49}")

	x.IntersectWith(z)
	checkValues(t, &x, "{9 42}")

	x.AddAll(138, 999)
	x.IntersectWith(z)
	checkValues(t, &x, "{9 42 138}")

	x.AddAll(6, 999)
	z.DifferenceWith(&x)
	checkValues(t, z, "{1 68 188}")

	z.AddAll(2, 6, 42)
	z.SymmetricDifference(&x)
	checkValues(t, z, "{1 2 9 68 138 188 999}")

	elems := z.Elems()
	eString := "[1 2 9 68 138 188 999]"
	if eString != fmt.Sprint(elems) {
		t.Fatalf("elems=%q, want %q", fmt.Sprint(elems), eString)
	}
}

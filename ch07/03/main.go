package main

import "fmt"

func main() {
	t := add(nil, 10)
	t = add(t, 20)
	t = add(t, 6)
	t = add(t, 1)
	t = add(t, 8)
	t = add(t, 4)
	t = add(t, 30)

	fmt.Println(t)
}

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}

	return fmt.Sprintf("{%d [%s] [%s]}", t.value, t.left.String(), t.right.String())
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

package main

type MapSet struct {
	m map[int]bool
}

func NewMapSet() *MapSet {
	return &MapSet{m: make(map[int]bool)}
}

func (s *MapSet) Add(x int) {
	s.m[x] = true
}

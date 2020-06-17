package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

var uintLength int = 32 << (^uint(0) >> 63)

// Len returns the number of elements
func (s *IntSet) Len() (length int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintLength; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}

	return
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintLength, uint(x%uintLength)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintLength, uint(x%uintLength)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds all values in list to the set
func (s *IntSet) AddAll(list ...int) {
	for _, v := range list {
		s.Add(v)
	}
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/uintLength, uint(x%uintLength)
	if word < len(s.words) {
		s.words[word] &= ^(1 << bit)
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	for i := 0; i < len(s.words); i++ {
		s.words[i] &= 0
	}
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var c IntSet
	c.words = make([]uint, len(s.words), len(s.words))
	for i := 0; i < len(s.words); i++ {
		c.words[i] = s.words[i]
	}

	return &c
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersect of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i >= len(t.words) {
			s.words[i] &= 0
		} else {
			s.words[i] &= t.words[i]
		}
	}
}

// DifferenceWith sets s to the difference of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i < len(t.words) {
			s.words[i] &= ^t.words[i]
		}
	}
}

// SymmetricDifference sets s to the set of the elements present in one set or the other but not both
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems returns a slice containing the elements of the set
func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < uintLength; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, uintLength*i+j)
			}
		}
	}

	return
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < uintLength; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintLength*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

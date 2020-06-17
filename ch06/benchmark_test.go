package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkIntSetAdd(b *testing.B) {
	rs := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rs)
	var s IntSet
	for i := 0; i < b.N; i++ {
		s.Add(r.Intn(10000))
	}
}

func BenchmarkMapSetAdd(b *testing.B) {
	rs := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rs)
	s := NewMapSet()
	for i := 0; i < b.N; i++ {
		s.Add(r.Intn(10000))
	}
}

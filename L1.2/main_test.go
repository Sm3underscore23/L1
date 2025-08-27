package main

import (
	"math/rand"
	"testing"
)

var array = rand.Perm(1000)

func BenchmarkWithAtomic(b *testing.B) {
	for range b.N {
		withAtomic(array)
	}
}

func BenchmarkWithMutex(b *testing.B) {
	for range b.N {
		withMutex(array)
	}

}

func BenchmarkWithChan(b *testing.B) {
	for range b.N {
		withChan(array)
	}
}

package main

import (
	"testing"
)

const order = 100 * 1000

var list StringList
var set StringSet

func BenchmarkListAdd(b *testing.B) {
	var sl StringList
	last := ""
	for i := 0; i < order; i++ {
		next := NextAlias(last)
		sl = append(sl, next)
		last = next
	}
	list = sl
}

func BenchmarkSetAdd(b *testing.B) {
	ss := make(StringSet)
	last := ""
	for i := 0; i < order; i++ {
		next := NextAlias(last)
		ss.Add(next)
		last = next
	}
	set = ss
}

func BenchmarkListUnite(b *testing.B) {
}

func BenchmarkSetUnite(b *testing.B) {
}

func BenchmarkListContainsHit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rndStr := RandStringRunes(2, "lower")
		list.Contains(rndStr)
	}
}

func BenchmarkSetContainsHit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rndStr := RandStringRunes(2, "lower")
		set.Contains(rndStr)
	}
}

func BenchmarkListContainsNonHit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rndStr := RandStringRunes(2, "upper")
		list.Contains(rndStr)
	}
}

func BenchmarkSetContainsNonHit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rndStr := RandStringRunes(2, "upper")
		set.Contains(rndStr)
	}
}

func BenchmarkListRemove(b *testing.B) {
}

func BenchmarkSetRemove(b *testing.B) {
}

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

// func BenchmarkListUnite(b *testing.B) {
// }

// func BenchmarkSetUnite(b *testing.B) {
// }

func BenchmarkListContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rndStr := RandStringRunes(3)
		list.Contains(rndStr)
	}
}

func BenchmarkSetContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rndStr := RandStringRunes(3)
		set.Contains(rndStr)
	}
}

// func BenchmarkListRemove(b *testing.B) {
// }

// func BenchmarkSetRemove(b *testing.B) {
// }

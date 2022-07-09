package main

import (
	"testing"
)

func BenchmarkListAdd(b *testing.B) {
	var sl StringList
	last := ""
	for i := 0; i < int(b.N); i++ {
		next := NextAlias(last)
		sl = append(sl, next)
		last = next
	}
}

func BenchmarkSetAdd(b *testing.B) {
	ss := make(StringSet)
	last := ""
	for i := 0; i < int(b.N); i++ {
		next := NextAlias(last)
		ss.Add(next)
		last = next
	}
}

// func BenchmarkListUnite(b *testing.B) {
// }

// func BenchmarkSetUnite(b *testing.B) {
// }

func BenchmarkListContains(b *testing.B) {
	sl := NewStringList(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rndStr := RandStringRunes(3)
		sl.Contains(rndStr)
	}
}

func BenchmarkSetContains(b *testing.B) {
	ss := NewStringSet(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rndStr := RandStringRunes(3)
		ss.Contains(rndStr)
	}
}

// func BenchmarkListRemove(b *testing.B) {
// }

// func BenchmarkSetRemove(b *testing.B) {
// }

package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

type StringList []string
type StringSet map[string]struct{}

var list StringList
var set StringSet

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NextAlias(last string) string {
	if last == "" {
		return "a"
	} else if last[len(last)-1] == 'z' {
		return last[:len(last)-1] + "aa"
	} else {
		return last[:len(last)-1] + string(last[len(last)-1]+1)
	}
}

func NewStringList(n int) (stringList StringList) {
	last := ""
	for i := 0; i < n; i++ {
		next := NextAlias(last)
		stringList = append(stringList, next)
		last = next
	}
	return
}

func NewStringSet(n int) StringSet {
	stringSet := make(StringSet, n)
	last := ""
	for i := 0; i < n; i++ {
		next := NextAlias(last)
		stringSet[next] = struct{}{}
		last = next
	}
	return stringSet
}

func NewStringUnions(n int) (StringList, StringSet) {
	stringList := make(StringList, n)
	stringSet := make(StringSet, n)
	last := ""
	for i := 0; i < n; i++ {
		next := NextAlias(last)
		stringList[i] = next
		stringSet[next] = struct{}{}
		last = next
	}
	return stringList, stringSet
}

func (sl StringList) Contains(str string) bool {
	for _, item := range sl {
		if item == str {
			return true
		}
	}
	return false
}

func (ss StringSet) Contains(str string) bool {
	_, ok := ss[str]
	return ok
}

func BenchmarkListContains(b *testing.B) {
	randomStr := RandStringRunes(3)
	b.ResetTimer()
	fmt.Printf("%s is contained in List: %t\n", randomStr, list.Contains(randomStr))
}

func BenchmarkSetContains(b *testing.B) {
	randomStr := RandStringRunes(3)
	b.ResetTimer()
	fmt.Printf("%s is contained in Set : %t\n", randomStr, set.Contains(randomStr))
}

func init() {
	// TO DO : Contains() always hit vs Contains() never hit
	list, set = NewStringUnions(500 * 1000)
	fmt.Printf("Length of List: %d, Length of Set: %d\n", len(list), len(set))
	fmt.Printf("Size  of  List: %d, Size  of  Set: %d\n", unsafe.Sizeof(list), unsafe.Sizeof(set))
}

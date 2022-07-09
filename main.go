package main

import (
	"fmt"
	"unsafe"
)

type StringList []string
type StringSet map[string]struct{}

func (ss *StringSet) Add(str string) *StringSet {
	(*ss)[str] = struct{}{}
	return ss
}

func (sl *StringList) Unite(sl2 StringList) *StringList {
	// TO dO
	return sl
}

func (ss *StringSet) Unite(ss2 StringSet) *StringSet {
	// TO dO
	return ss
}

func (sl StringList) Contains(str string) bool {
	for _, item := range sl {
		if item == str {
			return true
		}
	}
	return false
}

func (ss *StringSet) Contains(str string) bool {
	_, ok := (*ss)[str]
	return ok
}

func (sl StringList) Remove(str string) StringList {
	// TO DO
	return sl
}

func (ss *StringSet) Remove(str string) *StringSet {
	delete(*ss, str)
	return ss
}

func NewStringList(n int) StringList {
	sl := make(StringList, n)
	last := ""
	for i := 0; i < n; i++ {
		next := NextAlias(last)
		sl[i] = next
		last = next
	}
	return sl
}

func NewStringSet(n int) *StringSet {
	ss := make(StringSet, n)
	last := ""
	for i := 0; i < n; i++ {
		next := NextAlias(last)
		ss.Add(next)
		last = next
	}
	return &ss
}

func NewStringListSet(n int) (StringList, *StringSet) {
	sl := make(StringList, n)
	ss := make(StringSet, n)
	last := ""
	for i := 0; i < n; i++ {
		next := NextAlias(last)
		sl[i] = next
		ss.Add(next)
		last = next
	}
	return sl, &ss
}

func main() {
	list, set := NewStringListSet(100 * 1000)
	fmt.Printf("Length of List: %d, Length of Set: %d\n", len(list), len(*set))
	fmt.Printf("Size  of  List: %d, Size  of  Set: %d\n", unsafe.Sizeof(list), unsafe.Sizeof(*set))

	randomStr := RandStringRunes(3, "lower")
	fmt.Printf("%s is contained in List: %t\n", randomStr, list.Contains(randomStr))
	fmt.Printf("%s is contained in Set : %t\n", randomStr, set.Contains(randomStr))
}

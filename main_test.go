package main

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	nl0 := NewStringList(0)
	anl0 := StringList{}
	if !reflect.DeepEqual(nl0, anl0) {
		t.Error(nl0, anl0)
	}
	nl1 := NewStringList(1)
	anl1 := StringList{"a"}
	if !reflect.DeepEqual(nl1, anl1) {
		t.Error(nl1, anl1)
	}
	nl5 := NewStringList(5)
	anl5 := StringList{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(nl5, anl5) {
		t.Error(nl5, anl5)
	}
}

func TestNewSet(t *testing.T) {
	ns0 := NewStringSet(0)
	ans0 := StringSet{}
	if !reflect.DeepEqual(*ns0, ans0) {
		t.Error(ns0, ans0)
	}
	ns1 := NewStringSet(1)
	ans1 := make(StringSet, 1)
	ans1["a"] = struct{}{}
	if !reflect.DeepEqual(*ns1, ans1) {
		t.Error(ns1, ans1)
	}
	ns5 := NewStringSet(5)
	ans5 := make(StringSet, 5)
	ans5["a"] = struct{}{}
	ans5["b"] = struct{}{}
	ans5["c"] = struct{}{}
	ans5["d"] = struct{}{}
	ans5["e"] = struct{}{}
	if !reflect.DeepEqual(*ns5, ans5) {
		t.Error(ns5, ans5)
	}
}

func TestAppend(t *testing.T) {
	nl0 := []string{}
	nl0 = append(nl0, "a")
	nl1 := []string{"a"}
	if !reflect.DeepEqual(nl0, nl1) {
		t.Error(nl0, nl1)
	}

	nsl0 := StringList{}
	nsl0 = append(nsl0, "a")
	nsl1 := StringList{"a"}
	if !reflect.DeepEqual(nsl0, nsl1) {
		t.Error(nsl0, nsl1)
	}

}

func TestListAdd(t *testing.T) {
	nl0 := NewStringList(0)
	nl0 = append(nl0, "a")
	anl1 := StringList{"a"}
	if !reflect.DeepEqual(nl0, anl1) {
		t.Error(nl0, anl1)
	}

	nl1 := NewStringList(1)
	nl1 = append(nl1, "b")
	anl2 := StringList{"a", "b"}
	if !reflect.DeepEqual(nl1, anl2) {
		t.Error(nl1, anl2)
	}
}

func TestSetAdd(t *testing.T) {
	ns0 := NewStringSet(0)
	ns0.Add("a")
	ans1 := make(StringSet, 1)
	ans1["a"] = struct{}{}
	if !reflect.DeepEqual(*ns0, ans1) {
		t.Error(ns0, ans1)
	}
	ns1 := NewStringSet(1)
	ns1.Add("b")
	ans2 := make(StringSet, 2)
	ans2["a"] = struct{}{}
	ans2["b"] = struct{}{}
	if !reflect.DeepEqual(*ns1, ans2) {
		t.Error(ns1, ans2)
	}
}

func TestListContains(t *testing.T) {
	nl5 := NewStringList(5)
	if !nl5.Contains("a") {
		t.Error(nl5, "a")
	}
	if nl5.Contains("A") {
		t.Error(nl5, "A")
	}
	if nl5.Contains("") {
		t.Error(nl5, "")
	}
}

func TestSetContains(t *testing.T) {
	ns5 := NewStringSet(5)
	if !ns5.Contains("a") {
		t.Error(ns5, "a")
	}
	if ns5.Contains("A") {
		t.Error(ns5, "A")
	}
	if ns5.Contains("") {
		t.Error(ns5, "")
	}
}

func TestListReomve(t *testing.T) {
}

func TestSetReomve(t *testing.T) {
	ns1 := NewStringSet(1)
	ns1.Remove("b") // should be nothing
	ans1 := make(StringSet, 1)
	ans1["a"] = struct{}{}
	if !reflect.DeepEqual(*ns1, ans1) {
		t.Error(*ns1, ans1)
	}
	ns2 := NewStringSet(2)
	ns2.Remove("a")
	ans2 := make(StringSet, 1)
	ans2["b"] = struct{}{}
	if !reflect.DeepEqual(*ns2, ans2) && !ns1.Contains("b") {
		t.Error(*ns2, ans2)
	}
}

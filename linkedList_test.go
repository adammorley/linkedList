package linkedList

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

type Int int

func (i Int) LessThan(j interface{}) bool    { return i < j.(Int) }
func (i Int) GreaterThan(j interface{}) bool { return i > j.(Int) }
func (i Int) EqualTo(j interface{}) bool     { return i == j.(Int) }
func (i Int) String() string                 { return strconv.Itoa(int(i)) }

func TestString(t *testing.T) {
	l := New()
	l.Add(Int(5))
	l.Add(Int(6))
	if l.String() != "length: 2; h: -> 5 -> 6 :t" {
		t.Error("string format changed")
	}
}

func TestAdd(t *testing.T) {
	l := New()
	l.Add(Int(5))
	if l.Count(Int(5)) != 1 {
		t.Error("Inserted 5 but could not find it")
	}
	l.Add(Int(5))
	if l.Count(Int(5)) != 2 {
		t.Error("tested double add and no worky")
	}
}

func TestLength(t *testing.T) {
	l := New()
	l.Add(Int(1))
	l.Add(Int(2))
	l.Add(Int(3))
	if l.Length() != 3 {
		t.Error("length fails", l.Length())
	}
}

func TestCount(t *testing.T) {
	l := New()
	l.Add(Int(5))
	l.Add(Int(6))
	l.Add(Int(5))
	if l.Count(Int(5)) != 2 {
		t.Error("error counting 5")
	} else if l.Count(Int(6)) != 1 {
		t.Error("error counting 6")
	} else if l.Count(Int(7)) != 0 {
		t.Error("error counting 7")
	}
}

func TestDelete(t *testing.T) {
	l := New()
	l.Add(Int(5))
	l.Add(Int(6))
	l.Add(Int(6))
	l.Add(Int(7))
	if l.Delete(Int(4)) {
		t.Error("assert")
	}
	if !l.Delete(Int(5)) {
		t.Error("delete failes")
	}
	if l.Count(Int(5)) != 0 {
		t.Error("didn't delete")
	}
	if !l.Delete(Int(6)) {
		t.Error("delete failes")
	}
	if l.Count(Int(6)) != 0 {
		t.Error("deleted but then found 6")
	}
}

func TestMerge(t *testing.T) {
	l0 := New()
	l1 := New()
	l0.Add(Int(1))
	l0.Add(Int(4))
	l0.Add(Int(5))
	l1.Add(Int(2))
	l1.Add(Int(3))
	l1.Add(Int(6))
	l2 := new(LinkedList)
	merge(l2, l0, l1)
	ln := l2.head
	var i int = 1
	for ln != nil {
		if !ln.val.EqualTo(Int(i)) {
			t.Error("merge failure")
		}
		ln = ln.next
		i++
	}
}

func TestSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		l := New()
		n := randIntSlice(i)
		for _, v := range n {
			l.Add(Int(v))
		}
		l.Sort()
		ln := l.head
		sort.Sort(sort.IntSlice(n))
		j := 0
		for ln != nil {
			if !ln.val.EqualTo(Int(n[j])) {
				t.Error("problem", n, l)
			}
			j++
			ln = ln.next
		}
	}
}
func randIntSlice(v int) (r []int) {
	used := map[int]bool{}
	for i := 0; i < 10; {
		n := rand.Intn(100)
		if v%2 == 0 {
			r = append(r, n)
			i++
		} else if v%2 != 0 && !used[n] {
			used[n] = true
			r = append(r, n)
			i++
		}
	}
	return
}

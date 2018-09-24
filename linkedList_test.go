package linkedList

import "math/rand"
import "sort"
import "testing"

func TestString(t *testing.T) {
	l := New()
	l.Add(5)
	l.Add(6)
	if l.String() != "length: 2; h: -> 5 -> 6 :t" {
		t.Error("string format changed")
	}
}

func TestAdd(t *testing.T) {
	l := New()
	l.Add(5)
	if l.Count(5) != 1 {
		t.Error("Inserted 5 but could not find it")
	}
	l.Add(5)
	if l.Count(5) != 2 {
		t.Error("tested double add and no worky")
	}
}

func TestLength(t *testing.T) {
	l := New()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.Length() != 3 {
		t.Error("length fails", l.Length())
	}
}

func TestCount(t *testing.T) {
	l := New()
	l.Add(5)
	l.Add(6)
	l.Add(5)
	if l.Count(5) != 2 {
		t.Error("error counting 5")
	} else if l.Count(6) != 1 {
		t.Error("error counting 6")
	} else if l.Count(7) != 0 {
		t.Error("error counting 7")
	}
}

func TestDelete(t *testing.T) {
	l := New()
	l.Add(5)
	l.Add(6)
	l.Add(6)
	l.Add(7)
	if l.Delete(4) {
		t.Error("assert")
	}
	if !l.Delete(5) {
		t.Error("delete failes")
	}
	if l.Count(5) != 0 {
		t.Error("didn't delete")
	}
	if !l.Delete(6) {
		t.Error("delete failes")
	}
	if l.Count(6) != 0 {
		t.Error("deleted but then found 6")
	}
}

func TestMerge(t *testing.T) {
	l0 := New()
	l1 := New()
	l0.Add(1)
	l0.Add(4)
	l0.Add(5)
	l1.Add(2)
	l1.Add(3)
	l1.Add(6)
	l2 := new(LinkedList)
	merge(l2, l0, l1)
	ln := l2.head
	var i int = 1
	for ln != nil {
		if ln.val != i {
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
			l.Add(v)
		}
		l.Sort()
		ln := l.head
		sort.Sort(sort.IntSlice(n))
		j := 0
		for ln != nil {
			if ln.val != n[j] {
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
		if !used[n] && v%2 == 0 {
			used[n] = true
			r = append(r, n)
			i++
		} else {
			r = append(r, n)
			i++
		}
	}
	return
}

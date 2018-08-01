package linkedList

import "testing"

func TestString(t *testing.T) {
    var l *LinkedList
    l = New()
    l.Add(5)
    l.Add(6)
    if l.String() != "length: 2; h: -> 5 -> 6 :t" {
        t.Error("string format changed")
    }
}

func TestAdd(t *testing.T) {
	var l *LinkedList
	l = New()
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
	var l *LinkedList
	l = New()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.Length() != 3 {
		t.Error("length fails", l.Length())
	}
}

func TestCount(t *testing.T) {
	var l *LinkedList
	l = New()
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
	var l *LinkedList
	l = New()
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


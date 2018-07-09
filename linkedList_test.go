package linkedList

import "testing"

func TestAdd(t *testing.T) {
    var l *LinkedList
    l = New()
    l.Add(5)
    if l.Search(5) != 1 {
        t.Error("Inserted 5 but could not find it")
    }
    l.Add(5)
    if l.Search(5) != 2 {
        t.Error("tested double add and no worky")
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
    if ! l.Delete(5) {
        t.Error("delete failes")
    }
    if l.Search(5) != 0 {
        t.Error("didn't delete")
    }
    if ! l.Delete(6) {
        t.Error("delete failes")
    }
    if l.Search(6) != 0 {
        t.Error("deleted but then found 6")
    }
}

func TestLength(t *testing.T) {
    var l *LinkedList
    l = New()
    l.Add(1)
    l.Add(2)
    l.Add(3)
    if l.Length() != 3 {
        t.Error("length fails")
    }
}

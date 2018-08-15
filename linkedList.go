/*
   doubly-linked list
*/
package linkedList

import (
	"strconv"
	"strings"
)

type LinkedNode struct {
	prev *LinkedNode
	next *LinkedNode
	val  int64
}

type LinkedList struct {
	head   *LinkedNode
	tail   *LinkedNode
	length uint64
}

func New() *LinkedList {
	var l *LinkedList = new(LinkedList)
	l.head = nil
	l.tail = nil
	return l
}

/*
	Add to the list; goes on the back
*/
func (l *LinkedList) Add(v int64) {
	var ln *LinkedNode = new(LinkedNode)
	ln.val = v
	l.length++
	if l.head == nil {
		l.head = ln
		l.tail = ln
	} else {
		var i *LinkedNode = l.head
		for i.next != nil {
			i = i.next
		}
		i.next = ln
		ln.prev = i
		l.tail = ln
	}
	return
}

/*
	count the number of times a value appears in the list
*/
func (l *LinkedList) Count(v int64) (c uint64) {
	var ln *LinkedNode = l.head
	for ln != nil {
		if ln.val == v {
			c++
		}
		ln = ln.next
	}
	return
}

/*
   Delete removes all values equal to v in the linked
   list.  returns the count of the number of values
   deleted.
*/
func (l *LinkedList) Delete(v int64) bool {
	var b bool = false
	var t *LinkedNode
	var c *LinkedNode = l.head
	for c != nil {
		t = c.next
		if c.val == v {
			b = true
			if c.next == nil { // last element
				l.tail = c.prev
				c.prev.next = nil
			} else if c.prev == nil { // first element
				l.head = c.next
				c.next.prev = nil
			} else {
				c.prev.next = c.next
				c.next.prev = c.prev
			}
			l.length--
		}
		c = t
	}
	return b
}

func (l *LinkedList) Length() uint64 {
	return l.length
}

/*
   sort the list using a basic merge sort
*/
func (l *LinkedList) Sort() {

}

/*
   string-i-fy a linked list
*/
func (l *LinkedList) String() string {
	var b strings.Builder
	b.Grow(16)
	b.WriteString("length: ")
	b.WriteString(strconv.FormatUint(l.length, 10))
	b.WriteString("; h:")
	var ln *LinkedNode = l.head
	for ln != nil {
		b.WriteString(" -> ")
		b.WriteString(strconv.FormatInt(ln.val, 10))
		ln = ln.next
	}
	b.WriteString(" :t")
	return b.String()
}

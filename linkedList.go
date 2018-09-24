/*
   doubly-linked list
*/
package linkedList

import (
	"strconv"
	"strings"
)

/*
    consumers of this package must implement the following interfaces in order for the AVL tree to order the elements of the tree.  it's critical to use a type assertion; as a type assertion can cause a panic if the types don't match, optionally one can use a type-testing assertion (eg: i, ok := j.(TYPE) if !ok ...).  however, this would require a minor modification to the Interface of this package, as it would need to allow for error handling (eg: type mismatch)

type MyInt int

func (i MyInt) LessThan(j interface{}) bool {
    return i < j.(MyInt)
}

func (i MyInt) GreaterThan(j interface{}) bool {
    return i > j.(MyInt)
}

func (i MyInt) EqualTo(j interface{}) bool {
    return i == j.(MyInt)
}

func (i MyInt) String() string {
    return strconv.Itoa(int(i))
}
*/

type Interface interface {
	LessThan(j interface{}) bool
	GreaterThan(j interface{}) bool
	EqualTo(j interface{}) bool
	String() string
}

type LinkedNode struct {
	prev *LinkedNode
	next *LinkedNode
	val  Interface
}

type LinkedList struct {
	head   *LinkedNode
	tail   *LinkedNode
	length int
}

func New() (l *LinkedList) {
	l = new(LinkedList)
	l.head = nil
	l.tail = nil
	return
}

/*
	Add to the list; goes on the back
*/
func (l *LinkedList) Add(i Interface) {
	ln := new(LinkedNode)
	ln.val = i
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
func (l *LinkedList) Count(i Interface) (c int) {
	var ln *LinkedNode = l.head
	for ln != nil {
		if ln.val.EqualTo(i) {
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
func (l *LinkedList) Delete(i Interface) (b bool) {
	var t *LinkedNode
	var c *LinkedNode = l.head
	for c != nil {
		t = c.next
		if c.val.EqualTo(i) {
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
	return
}

func (l *LinkedList) Length() int {
	return l.length
}

/*
   sort the list using a basic merge sort

   technically, the ints could be inserted into an avlTree and read out in order
   additionally, if the values were read out into a slice (or if the implementation were modified to be
   a slice), the sort package could sort this very readily.

   additionally, using the heap package would also be an option (eg: a min heap)

   but for the purposes of this code, which is about learning, let's do a merge sort.

   a merge sort is when the list is split into smaller and smaller lists, until the list is separated into
   pairs.  the pairs are sorted, and then pairs are merged together recursively.  this allows the list
   to end up sorted at the end
*/
func (l *LinkedList) Sort() {
	length := l.Length()
	if length < 2 {
		return
	}
	left := New()
	right := New()
	left.head = l.head
	right.tail = l.tail

	midpoint := length / 2
	var cur *LinkedNode = l.head
	for i := 1; i <= midpoint; i++ {
		if i == midpoint {
			right.head = cur.next
			right.head.prev = nil
			left.tail = cur
			left.tail.next = nil
			break
		}
		cur = cur.next
	}
	left.length = midpoint
	right.length = length - midpoint

	left.Sort()
	right.Sort()
	merge(l, left, right)
	return
}
func merge(list, left, right *LinkedList) {
	ln, rn := left.head, right.head
	head := true
	var cur *LinkedNode
	// go through the two lists & compare
	for ln != nil && rn != nil {
		if ln.val.LessThan(rn.val) || ln.val.EqualTo(rn.val) {
			if !head {
				ln.prev = cur
				cur.next = ln
			}
			cur = ln
			ln = ln.next
		} else if ln.val.GreaterThan(rn.val) {
			if !head {
				rn.prev = cur
				cur.next = rn
			}
			cur = rn
			rn = rn.next
		}
		if head {
			list.head = cur
			head = false
		}
	}

	// check for leftovers
	for ln != nil {
		ln.prev = cur
		cur.next = ln
		cur = ln
		ln = ln.next
	}
	for rn != nil {
		rn.prev = cur
		cur.next = rn
		cur = rn
		rn = rn.next
	}

	// update the tail
	list.tail = cur
	list.tail.next = nil
	list.length = left.Length() + right.Length()

	return
}

/*
   string-i-fy a linked list
*/
func (l *LinkedList) String() string {
	var b strings.Builder
	b.Grow(16)
	b.WriteString("length: ")
	b.WriteString(strconv.Itoa(int(l.length)))
	b.WriteString("; h:")
	var ln *LinkedNode = l.head
	for ln != nil {
		b.WriteString(" -> ")
		b.WriteString(ln.val.String())
		ln = ln.next
	}
	b.WriteString(" :t")
	return b.String()
}

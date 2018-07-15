/*
   doubly-linked list
*/
package linkedList

type LinkedNode struct {
	prev *LinkedNode
	next *LinkedNode
	val  int
}

type LinkedList struct {
	head *LinkedNode
	tail *LinkedNode
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
func (l *LinkedList) Add(v int) {
	if l.head == nil {
		var ln *LinkedNode = new(LinkedNode)
		ln.val = v
		l.head = ln
		l.tail = ln
		return
	}
	var ln *LinkedNode = l.head
	for {
		if ln.next != nil {
			ln = ln.next
		} else {
			ln.next = new(LinkedNode)
			ln.next.prev = ln
			ln.next.val = v
			l.tail = ln.next
			return
		}
	}
}

/*
	count the number of times a value appears in the list
*/
func (l *LinkedList) Count(v int) (c int) {
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
func (l *LinkedList) Delete(v int) bool {
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
		}
		c = t
	}
	return b
}

func (l *LinkedList) Length() uint {
	var t *LinkedNode = l.head
	if t == nil {
		return 0
	}
	var c uint = 0
	for {
		c++
		if t.next != nil {
			t = t.next
		} else {
			return c
		}
	}
}

/*
   Search searches for the value v and returns the count stored in the linked list
*/
func (l *LinkedList) Search(v int) uint {
	if l.head == nil {
		return 0
	}
	var c uint = 0
	var ln *LinkedNode = l.head
	for {
		if ln.val == v {
			c++
		}
		if ln.next == nil {
			return c
		} else {
			ln = ln.next
		}
	}
}

/*
   string-i-fy a linked list

   this needs improvement to ensure that the data does not overflow when the list is big (there must be some mechanism in golang to re-call the stringer function with a continuation token; if not, uh-oh)
*/
func (l *LinkedList) String() string {
	var c uint = l.Length()
	if c == 0 {
		return ""
	}
	var s uint = c * (19 + 2)
	var r = make([]rune, s, s)
	var ln *LinkedNode = l.head
	for {
	}
}

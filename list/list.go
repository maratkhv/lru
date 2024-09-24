package list

type Element[V any] struct {
	next, prev *Element[V]
	parentList *List[V]
	Value      V
}

type List[V any] struct {
	head   Element[V]
	length int64
}

// New returns new instance of list
func New[V any]() *List[V] {
	l := &List[V]{}
	l.init()
	return l
}

// init inits new list if not already
func (l *List[V]) init() {
	if l.head.next == nil {
		l.head.next = &l.head
		l.head.prev = &l.head
		l.length = 0
	}
}

// Front returns first element
func (l *List[V]) Front() *Element[V] {
	if l.Len() == 0 {
		return nil
	}
	return l.head.next
}

// Back returns last element
func (l *List[V]) Back() *Element[V] {
	if l.Len() == 0 {
		return nil
	}
	return l.head.prev
}

// Len returns number of elements in a list
func (l *List[V]) Len() int64 {
	return l.length
}

// insert inserts value v after at
func (l *List[V]) insert(v V, at *Element[V]) (e *Element[V]) {
	e = &Element[V]{
		Value:      v,
		next:       at.next,
		prev:       at,
		parentList: l,
	}
	e.prev.next = e
	e.next.prev = e

	l.length++

	return e
}

// PushBack inserts value v to the end of list
func (l *List[V]) PushBack(v V) *Element[V] {
	return l.insert(v, l.head.prev)
}

// PushBack inserts value v to the front of list
func (l *List[V]) PushFront(v V) *Element[V] {
	return l.insert(v, l.head.next)
}

// InsertAfter inserts value v after at
func (l *List[V]) InsertAfter(v V, at *Element[V]) *Element[V] {
	return l.insert(v, at)
}

// InsertBefore inserts value v after at
func (l *List[V]) InsertBefore(v V, at *Element[V]) *Element[V] {
	return l.insert(v, at.prev)
}

// Pop removes element from the list and returns it even if it was not in the list
func (l *List[V]) Pop(e *Element[V]) *Element[V] {
	if e != nil && e.parentList == l {
		e.next.prev = e.prev
		e.prev.next = e.next
		l.length--
	}

	return e
}

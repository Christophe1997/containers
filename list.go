package containers

import "iter"

type LinkedList[E Elem] interface {
	SizedSeq[E]
	Front() E
	Back() E
	PushFront(e E) LinkedList[E]
	PushBack(e E) LinkedList[E]
	RemoveFront() LinkedList[E]
	RemoveBack() LinkedList[E]
}

func NewLinkedList[E Elem]() LinkedList[E] {
	return newLinkedList[E]()
}

func newLinkedList[E Elem]() *linkedList[E] {
	head := &node[E]{}
	tail := &node[E]{}
	head.prev = tail
	head.next = tail
	tail.next = head
	tail.prev = head
	return &linkedList[E]{
		head: head,
		tail: tail,
		size: 0,
	}
}

type node[E any] struct {
	next, prev *node[E]
	val        E
}

type linkedList[E Elem] struct {
	head *node[E]
	tail *node[E]
	size int
}

func (ls *linkedList[E]) Size() int {
	return ls.size
}

func (ls *linkedList[E]) All() iter.Seq2[int, E] {
	return ConvertSeq2(ls.Values())
}

func (ls *linkedList[E]) Values() iter.Seq[E] {
	return func(yield func(E) bool) {
		cur := ls.head.next
		for cur != ls.tail {
			if !yield(cur.val) {
				return
			}
			cur = cur.next
		}
	}
}

func (ls *linkedList[E]) Front() E {
	if ls.IsEmpty() {
		panic("Empty LinkedList")
	}
	return ls.head.next.val
}

func (ls *linkedList[E]) Back() E {
	if ls.IsEmpty() {
		panic("Empty LinkedList")
	}
	return ls.tail.prev.val
}

func (ls *linkedList[E]) PushFront(e E) LinkedList[E] {
	ls.insertAfter(ls.head, e)
	return ls
}

func (ls *linkedList[E]) PushBack(e E) LinkedList[E] {
	ls.insertBefore(ls.tail, e)
	return ls
}

func (ls *linkedList[E]) RemoveFront() LinkedList[E] {
	if ls.IsEmpty() {
		panic("Empty LinkedList")
	} else {
		ls.remove(ls.head.next)
		return ls
	}
}

func (ls *linkedList[E]) RemoveBack() LinkedList[E] {
	if ls.IsEmpty() {
		panic("Empty LinkedList")
	} else {
		ls.remove(ls.tail.prev)
		return ls
	}
}

func (ls *linkedList[E]) IsEmpty() bool {
	return ls.Size() == 0
}

func (ls *linkedList[E]) insertAfter(target *node[E], e E) {
	elem := &node[E]{val: e}
	elem.prev = target
	elem.next = target.next
	elem.prev.next = elem
	elem.next.prev = elem
	ls.size++
}

func (ls *linkedList[E]) insertBefore(target *node[E], e E) {
	elem := &node[E]{val: e}
	elem.next = target
	elem.prev = target.prev
	elem.next.prev = elem
	elem.prev.next = elem
	ls.size++
}

func (ls *linkedList[E]) remove(node *node[E]) {
	if ls.size > 0 {
		node.prev.next = node.next
		node.next.prev = node.prev
		node.prev, node.next = nil, nil
		ls.size--
	}
}

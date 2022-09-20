package containers

import (
	"fmt"
	"strings"
)

type List[E comparable] interface {
	Collection[E]
	fmt.Stringer
	Index(e E) (int, bool)
	IndexWith(cmp func(lhs, rhs E) bool, e E) (int, bool)
	Get(idx int) E
	Set(idx int, e E)
	ToArray() []E
}

type LinkedList[E comparable] interface {
	List[E]
	Front() E
	Back() E
	PushFront(e E) LinkedList[E]
	PushBack(e E) LinkedList[E]
	RemoveFront() LinkedList[E]
	RemoveBack() LinkedList[E]
}

func NewLinkedList[E comparable]() LinkedList[E] {
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

type linkedList[E comparable] struct {
	head *node[E]
	tail *node[E]
	size int
}

type linkedListIterator[E comparable] struct {
	data *linkedList[E]
	cur  *node[E]
}

func (it *linkedListIterator[E]) HasNext() bool {
	return it.cur.next != it.data.tail
}

func (it *linkedListIterator[E]) Next() E {
	if it.HasNext() {
		it.cur = it.cur.next
		return it.cur.val
	} else {
		panic("At end of Iterator")
	}
}

func (ls *linkedList[E]) ForEach(do func(E)) {
	it := ls.Iterator()
	for it.HasNext() {
		do(it.Next())
	}
}

func (ls *linkedList[E]) Iterator() Iterator[E] {
	return &linkedListIterator[E]{
		data: ls,
		cur:  ls.head,
	}
}

func (ls *linkedList[E]) Add(e ...E) {
	for _, v := range e {
		ls.PushBack(v)
	}
}

func (ls *linkedList[E]) Contains(e E) bool {
	_, has := ls.Index(e)
	return has
}

func (ls *linkedList[E]) ContainsWith(cmp func(lhs E, rhs E) bool, e E) bool {
	_, has := ls.IndexWith(cmp, e)
	return has
}

func (ls *linkedList[E]) Remove(e ...E) {
	if ls.IsEmpty() {
		panic("Empty LinkedList")
	}
	for _, v := range e {
		cur := ls.head.next
		for cur != ls.tail {
			if cur.val == v {
				temp := cur
				cur = cur.next
				ls.remove(temp)
			}
		}
	}
}

func (ls *linkedList[E]) RemoveWith(cmp func(lhs, rhs E) bool, e ...E) {
	if ls.IsEmpty() {
		panic("Empty LinkedList")
	}
	for _, v := range e {
		cur := ls.head.next
		for cur != ls.tail {
			if cmp(cur.val, v) {
				temp := cur
				cur = cur.next
				ls.remove(temp)
			}
		}
	}
}

func (ls *linkedList[E]) String() string {
	res := Fmap(func(e E) string {
		return fmt.Sprintf("%+v", e)
	})(ls)
	arr := res.ToArray()
	return fmt.Sprintf("[%s]", strings.Join(arr, ", "))
}

func (ls *linkedList[E]) Index(e E) (int, bool) {
	return ls.IndexWith(func(lhs, rhs E) bool { return lhs == rhs }, e)
}

func (ls *linkedList[E]) IndexWith(cmp func(lhs E, rhs E) bool, e E) (int, bool) {
	cur := ls.head.next
	res := 0
	for cur != ls.tail {
		if cmp(cur.val, e) {
			return res, true
		}
		cur = cur.next
		res++
	}
	return -1, false
}

func (ls *linkedList[E]) Get(idx int) E {
	if idx >= 0 && idx < ls.Size() {
		cur := ls.head.next
		for cur != ls.tail {
			if idx == 0 {
				return cur.val
			}
			cur = cur.next
			idx--
		}
	}
	panic("Out of bound")
}

func (ls *linkedList[E]) Set(idx int, e E) {
	if idx >= 0 && idx < ls.Size() {
		cur := ls.head.next
		for cur != ls.tail {
			if idx == 0 {
				cur.val = e
			}
			cur = cur.next
			idx--
		}
	}
	panic("Out of bound")
}

func (ls *linkedList[E]) ToArray() []E {
	res := make([]E, ls.Size())
	idx := 0
	ls.ForEach(func(e E) {
		res[idx] = e
		idx++
	})
	return res
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

func (ls *linkedList[E]) Size() int {
	return ls.size
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
	if ls.Size() > 0 {
		node.prev.next = node.next
		node.next.prev = node.prev
		node.prev, node.next = nil, nil
		ls.size--
	}
}

package containers

import "iter"

type Stack[E Elem] interface {
	SizedSeq[E]
	Pop() (Stack[E], E)
	Push(e E) Stack[E]
	Peek() E
}

type stack[E Elem] struct {
	data *linkedList[E]
}

func (s *stack[E]) All() (seq iter.Seq2[int, E]) {
	return ConvertSeq2(s.Values())
}

func (s *stack[E]) Values() (seq iter.Seq[E]) {
	return func(yield func(E) bool) {
		cur := s.data.tail.prev
		for cur != s.data.head {
			if !yield(cur.val) {
				return
			}
			cur = cur.prev
		}
	}
}

func NewStack[E Elem]() Stack[E] {
	return &stack[E]{data: newLinkedList[E]()}
}

func (s *stack[E]) Pop() (Stack[E], E) {
	if s.Size() == 0 {
		panic("Empty Stack")
	} else {
		res := s.Peek()
		s.data.RemoveBack()
		return s, res
	}
}

func (s *stack[E]) Push(e E) Stack[E] {
	s.data.PushBack(e)
	return s
}

func (s *stack[E]) Peek() E {
	return s.data.Back()
}

func (s *stack[E]) Size() int {
	return s.data.Size()
}

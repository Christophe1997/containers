package containers

import "iter"

type Queue[E any] interface {
	SizedSeq[E]
	Enqueue(e E) Queue[E]
	Dequeue() (Queue[E], E)
	Peek() E
}

type queue[E Elem] struct {
	data *linkedList[E]
}

func (q *queue[E]) All() (seq iter.Seq2[int, E]) {
	return q.data.All()
}

func (q *queue[E]) Values() (seq iter.Seq[E]) {
	return q.data.Values()
}

func NewQueue[E Elem]() Queue[E] {
	return &queue[E]{data: newLinkedList[E]()}
}

func (q *queue[E]) Size() int {
	return q.data.Size()
}

func (q *queue[E]) Enqueue(e E) Queue[E] {
	q.data.PushBack(e)
	return q
}

func (q *queue[E]) Dequeue() (Queue[E], E) {
	res := q.Peek()
	q.data.RemoveFront()
	return q, res
}

func (q *queue[E]) Peek() E {
	return q.data.Front()
}

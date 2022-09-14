package containers

type Queue[E comparable] interface {
	Sized

	Enqueue(e E) Queue[E]
	Dequeue() (Queue[E], E)
	Peek() E

	ToList() List[E]
}

type queue[E comparable] struct {
	data LinkedList[E]
}

func NewQueue[E comparable]() Queue[E] {
	return &queue[E]{data: NewLinkedList[E]()}
}

func (q *queue[E]) Size() int {
	return q.data.Size()
}

func (q *queue[E]) IsEmpty() bool {
	return q.data.IsEmpty()
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

func (q *queue[E]) ToList() List[E] {
	res := NewLinkedList[E]()
	Copy[E](res, q.data)
	return res
}
